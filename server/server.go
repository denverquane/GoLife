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

const DEBUG_BROADCAST_NON_REGISTERED = true

var upgrader = websocket.Upgrader{} // use default options

type Player struct {
	name  string
	color uint32
}

var clients = make(map[*websocket.Conn]Player)

var SimulationChannel = make(chan simulation.SimulatorMessage)

var addr = flag.String("addr", ":5000", "http service address")
var RleMap = make(map[string]simulation.RLE)

func main() {
	rle, err := simulation.LoadRLE("./data/glider.rle")
	if err != nil {
		log.Fatal(err)
	} else {
		RleMap["glider"] = rle
	}

	rle, err = simulation.LoadRLE("./data/pufferfish.rle")
	if err != nil {
		log.Fatal(err)
	} else {
		RleMap["pufferfish"] = rle
	}

	flag.Parse()
	log.SetFlags(0)

	Run(addr)
}

func Run(addr *string) {
	go simulationWorker(60, SimulationChannel)

	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var GlobalWorld simulation.World

func simulationWorker(targetFps int64, msgChan <-chan simulation.SimulatorMessage) {
	msPerFrame := (1.0 / float64(targetFps)) * 1000.0
	GlobalWorld = simulation.NewConwayWorld(200, 123)
	GlobalWorld.PlaceRLEAtCoords(RleMap["glider"], 0, 0, simulation.ALIVE_FULL)

	GlobalWorld.PlaceRLEAtCoords(RleMap["pufferfish"], 100, 150, simulation.ALIVE_FULL)
	paused := false
	for {
		select {
		case msg := <-msgChan:
			switch msg.Type {
			case simulation.TOGGLE_PAUSE:
				paused = !paused
			case simulation.MARK_CELL:
				if paused {
					GlobalWorld.MarkAliveColor(msg.Y, msg.X, msg.Color)
				}
			}
		default:
			if !paused && len(clients) > 0 {
				oldT := time.Now().UnixNano()
				GlobalWorld.Tick()
				//TODO send message to dedicated worker to send the status probably?
				//Consider race condition of message being received AFTER another tick...
				broadcastWorld(&GlobalWorld, false)
				//log.Print(GlobalWorld.ToString())
				tickMs := float64(time.Now().UnixNano()-oldT) / NS_PER_MS
				//log.Printf("%fms to tick; sleeping %fms\n", tickMs, msPerFrame - tickMs)
				//time.Sleep(time.Millisecond * 500)
				time.Sleep(time.Duration(NS_PER_MS * (msPerFrame - tickMs)))
			} else {
				broadcastWorld(&GlobalWorld, true)
				//log.Println("Simulation is paused; sleeping for 1000ms")
				time.Sleep(time.Millisecond * 50)
			}
		}
	}
}

func sendFirstWorldMessage(client *websocket.Conn, world *simulation.World) {
	marshalled, err := world.ToFullProtoBytes()

	if err != nil {
		log.Printf("Error in marshalling world: %s\n", err)
	} else {
		err := client.WriteMessage(websocket.BinaryMessage, marshalled)
		if err != nil {
			log.Println(err)
		}
	}
}

func broadcastWorld(world *simulation.World, paused bool) {
	marshalled, err := world.ToMinProtoBytes(paused)
	if err != nil {
		log.Printf("Error in marshalling world: %s\n", err)
	} else {
		for client, player := range clients {
			if player.name != "" || DEBUG_BROADCAST_NON_REGISTERED {
				err := client.WriteMessage(websocket.BinaryMessage, marshalled)
				if err != nil {
					log.Println(err)
				}
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
					//TODO verify name/color aren't taken
					log.Printf("Registering %s\n", regMsg.Name)
					clients[c] = Player{name: regMsg.Name, color: regMsg.Color}
					err := c.WriteMessage(websocket.BinaryMessage, data)
					if err != nil {
						log.Printf("Error echoing registration message to %s: %s\n", regMsg.Name, err)
					}
					broadcastPlayers()
					sendFirstWorldMessage(c, &GlobalWorld)
				}
			case message.MessageType_COMMAND:
				cmdMsg := message.Command{}
				err := proto.Unmarshal(msg.Content, &cmdMsg)
				if err != nil {
					log.Println(err)
				} else {
					switch cmdMsg.Type {
					case message.CommandType_TOGGLE_PAUSE:
						//TODO when more players are online, we prob want a voting mechanism for pausing
						log.Println("Sending toggle pause to channel")
						SimulationChannel <- simulation.SimulatorMessage{Type: simulation.TOGGLE_PAUSE}
					case message.CommandType_MARK_CELL:
						player := clients[c]
						SimulationChannel <- simulation.SimulatorMessage{
							Type:  simulation.MARK_CELL,
							X:     cmdMsg.X,
							Y:     cmdMsg.Y,
							Color: player.color,
						}
						log.Printf("Marking cell at (%d, %d) with color %32b", cmdMsg.X, cmdMsg.Y, player.color)
					}
				}
			default:
				log.Printf("Received non-recognized message of type %d with content: %s", msg.Type, msg.Content)
			}
		}
	}
}
