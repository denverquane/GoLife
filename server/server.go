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

var upgrader = websocket.Upgrader{} // use default options

var clients = make(map[*websocket.Conn]bool)

var addr = flag.String("addr", ":5000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	Run(addr)
}

func Run(addr *string) {
	go simulationWorker()

	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func simulationWorker() {
	conway := simulation.NewConwayWorld(400, 400)
	conway.MakeGliderGun(0, 0)
	for {
		conway.Tick()
		broadcastWorld(conway)
		//log.Print(conway.ToString())
		time.Sleep(time.Millisecond * 10)
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
	clients[c] = true

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
		err = proto.Unmarshal(data, &msg)
		if err != nil {
			log.Printf("Encountered error unmarshalling message: %s\n", err)
		} else {
			switch msg.Type {
			case message.MessageType_REGISTER:
				regMsg := message.RegisterName{}
				err := proto.Unmarshal(msg.Content, &regMsg)
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("Registering %s\n", regMsg.Name)
				}
			default:
				log.Printf("Received non-recognized message of type %d with content: %s", msg.Type, msg.Content)
			}
		}
	}
}
