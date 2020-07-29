package main

import (
	"flag"
	"github.com/denverquane/golife/proto/message"
	"github.com/denverquane/golife/simulation"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"time"
)

const NS_PER_MS = 1_000_000.0

const DEBUG_BROADCAST_NON_REGISTERED = false

var upgrader = websocket.Upgrader{} // use default options

type Player struct {
	name  string
	color uint32
}

var clients = make(map[*websocket.Conn]Player)

var addr = flag.String("addr", ":5000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	Run(addr)
}

func Run(addr *string) {
	go simulationWorker(60)

	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func simulationWorker(targetFps int64) {
	msPerFrame := (1.0 / float64(targetFps)) * 1000.0
	conway := simulation.NewConwayWorld(400, 400)
	conway.MakeGliderGun(0, 0)
	for {
		oldT := time.Now().UnixNano()
		conway.Tick()
		broadcastWorld(conway)
		tickMs := float64(time.Now().UnixNano()-oldT) / NS_PER_MS
		//log.Printf("%fms to tick; sleeping %fms\n", tickMs, msPerFrame - tickMs)
		time.Sleep(time.Duration(NS_PER_MS * (msPerFrame - tickMs)))
	}
}

func broadcastWorld(world simulation.World) {
	height, width := world.GetDims()
	worldMsg := message.WorldData{
		Width:  height,
		Height: width,
		Data:   world.GetFlattenedData(),
		Tick:   world.GetTick(),
	}
	worldMsgMarshalled, err := proto.Marshal(&worldMsg)
	if err != nil {
		log.Println(err)
		return
	}
	msg := message.Message{
		Type:    message.MessageType_WORLD_DATA,
		Content: worldMsgMarshalled,
	}
	marshalled, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}
	for client, player := range clients {
		if player.name != "" || DEBUG_BROADCAST_NON_REGISTERED {
			err := client.WriteMessage(websocket.BinaryMessage, marshalled)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func broadcastPlayers() {
	players := message.Players{}
	for _, v := range clients {
		players.Players = append(players.Players, &message.Player{
			Name:  v.name,
			Color: v.color,
		})
	}
	playersMarshalled, err := proto.Marshal(&players)
	if err != nil {
		log.Println(err)
		return
	}
	msg := message.Message{
		Type:    message.MessageType_PLAYERS,
		Content: playersMarshalled,
	}
	marshalled, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}
	for client := range clients {
		err := client.WriteMessage(websocket.BinaryMessage, marshalled)
		if err != nil {
			log.Println(err)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//TODO security, fix this once deployed
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	clients[c] = Player{
		name:  "",
		color: 0,
	}

	c.SetCloseHandler(func(code int, text string) error {
		log.Printf("Client disconnected with code %d and text: %s", code, text)
		delete(clients, c)
		return nil
	})

	defer c.Close()

	log.Printf("Client has connected with local addr %s and remote %s\n",
		c.LocalAddr().String(), c.RemoteAddr().String())

	for {
		msg := message.Message{}
		_, data, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		err = proto.Unmarshal(data, &msg)
		if err != nil {
			log.Printf("Encountered error unmarshalling message: %s\n", err)
		} else {
			switch msg.Type {
			case message.MessageType_REGISTER:
				regMsg := message.Player{}
				err := proto.Unmarshal(msg.Content, &regMsg)
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("Registering %s\n", regMsg.Name)
					clients[c] = Player{name: regMsg.Name, color: 10}
					err := c.WriteMessage(websocket.BinaryMessage, data)
					if err != nil {
						log.Printf("Error echoing registration message to %s: %s\n", regMsg.Name, err)
					}
					broadcastPlayers()
				}
			default:
				log.Printf("Received non-recognized message of type %d with content: %s", msg.Type, msg.Content)
			}
		}
	}
}
