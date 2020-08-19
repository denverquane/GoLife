package main

import (
	"flag"
	"github.com/denverquane/golife/proto/message"
	"github.com/denverquane/golife/simulation"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const NS_PER_MS = 1_000_000.0

const DEBUG_BROADCAST_NON_REGISTERED = true

const WORLD_DIM = 1000

var upgrader = websocket.Upgrader{} // use default options

type Player struct {
	name  string
	color uint32
}

var clients = make(map[*websocket.Conn]Player)
var clientsLock = sync.Mutex{}

type BroadcastType int

type BroadcastMsg struct {
	Btype  BroadcastType
	Paused bool
	Client *websocket.Conn
}

const (
	PLAYERS    BroadcastType = 0
	WORLD      BroadcastType = 1
	FIRST_DATA BroadcastType = 3
)

var SimulationChannel = make(chan simulation.SimulatorMessage)
var BroadcastChannel = make(chan BroadcastMsg)

var addr = flag.String("addr", ":5000", "http service address")
var RleMap = make(map[string]simulation.RLE)

func main() {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range files {
		if strings.HasSuffix(v.Name(), ".rle") {
			rle, err := simulation.LoadRLE("./data/" + v.Name())
			if err != nil {
				log.Println(err)
			} else {
				split := strings.Split(v.Name(), ".")
				RleMap[split[0]] = rle
			}
		}

	}

	flag.Parse()
	log.SetFlags(0)

	Run(addr)
}

func Run(addr *string) {
	var GlobalWorld simulation.World
	GlobalWorld = simulation.NewConwayWorld(WORLD_DIM, WORLD_DIM)
	go simulationWorker(&GlobalWorld, 60, SimulationChannel)
	go broadcastWorker(&GlobalWorld, BroadcastChannel)

	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func simulationWorker(world *simulation.World, targetFps int64, msgChan <-chan simulation.SimulatorMessage) {
	msPerFrame := (1.0 / float64(targetFps)) * 1000.0

	world.PlaceRLEAtCoords(RleMap["glider"], 0, 0, simulation.ALIVE_FULL)

	//GlobalWorld.PlaceRLEAtCoords(RleMap["pufferfish"], 100, 150, simulation.ALIVE_FULL)
	paused := false
	for {
		select {
		case msg := <-msgChan:
			switch msg.Type {
			case simulation.TOGGLE_PAUSE:
				paused = !paused
			case simulation.MARK_CELL:
				if paused {
					world.MarkAliveColor(msg.Y, msg.X, msg.Color)
				}
			case simulation.PLACE_RLE:
				if paused {
					for name, rle := range RleMap {
						if name == msg.Info {
							world.PlaceRLEAtCoords(rle, msg.Y, msg.X, msg.Color)
						}
					}
				}
			case simulation.CLEAR_BOARD:
				world.Clear()
			}
		default:
			clientsLock.Lock()
			numClients := len(clients)
			clientsLock.Unlock()
			if !paused && numClients > 0 {
				oldT := time.Now().UnixNano()
				world.Tick(true)

				//Consider race condition of message being received AFTER another tick...
				BroadcastChannel <- BroadcastMsg{
					Btype:  WORLD,
					Paused: false,
				}
				//log.Print(GlobalWorld.ToString())
				tickMs := float64(time.Now().UnixNano()-oldT) / NS_PER_MS
				//log.Printf("%fms to tick; sleeping %fms\n", tickMs, msPerFrame - tickMs)
				//time.Sleep(time.Millisecond * 500)
				time.Sleep(time.Duration(NS_PER_MS * (msPerFrame - tickMs)))
			} else {
				BroadcastChannel <- BroadcastMsg{
					Btype:  WORLD,
					Paused: true,
				}
				//log.Println("Simulation is paused; sleeping for 1000ms")
				time.Sleep(time.Millisecond * 50)
			}
		}
	}
}

func broadcastWorker(world *simulation.World, broadcasts <-chan BroadcastMsg) {
	for {
		select {
		case msg := <-broadcasts:
			switch msg.Btype {
			case PLAYERS:
				broadcastPlayers()
			case WORLD:
				broadcastWorld(world, msg.Paused)
			case FIRST_DATA:
				sendFirstWorldMessage(msg.Client, world)
				sendRLEs(msg.Client)
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

func sendRLEs(client *websocket.Conn) {
	rlesBytes := simulation.ToRleBytes(RleMap)
	msg := message.Message{
		Type:    message.MessageType_RLE_OPTIONS,
		Content: rlesBytes,
	}
	msgBytes, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}

	err = client.WriteMessage(websocket.BinaryMessage, msgBytes)
	if err != nil {
		log.Println(err)
		delete(clients, client)
	}
}

func broadcastWorld(world *simulation.World, paused bool) {
	marshalled, err := world.ToMinProtoBytes(paused)
	if err != nil {
		log.Printf("Error in marshalling world: %s\n", err)
	} else {
		clientsLock.Lock()
		for client, player := range clients {
			if player.name != "" || DEBUG_BROADCAST_NON_REGISTERED {
				err := client.WriteMessage(websocket.BinaryMessage, marshalled)
				if err != nil {
					log.Println(err)
					delete(clients, client)
				}
			}
		}
		clientsLock.Unlock()
	}
}

func broadcastPlayers() {
	serverData := message.ServerData{}
	clientsLock.Lock()
	for _, v := range clients {
		serverData.Players = append(serverData.Players, &message.Player{
			Name:  v.name,
			Color: v.color,
		})
	}
	clientsLock.Unlock()
	playersMarshalled, err := proto.Marshal(&serverData)
	if err != nil {
		log.Println(err)
		return
	}
	msg := message.Message{
		Type:    message.MessageType_SERVER_DATA,
		Content: playersMarshalled,
	}
	marshalled, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}
	clientsLock.Lock()
	for client := range clients {
		err := client.WriteMessage(websocket.BinaryMessage, marshalled)
		if err != nil {
			log.Println(err)
			delete(clients, client)
		}
	}
	clientsLock.Unlock()
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
	clientsLock.Lock()
	clients[c] = Player{
		name:  "",
		color: 0,
	}
	clientsLock.Unlock()

	c.SetCloseHandler(func(code int, text string) error {
		log.Printf("Client disconnected with code %d and text: %s", code, text)
		clientsLock.Lock()
		delete(clients, c)
		clientsLock.Unlock()
		BroadcastChannel <- BroadcastMsg{
			Btype: PLAYERS,
		}
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
					clientsLock.Lock()
					clients[c] = Player{name: regMsg.Name, color: regMsg.Color}
					err := c.WriteMessage(websocket.BinaryMessage, data)
					clientsLock.Unlock()
					if err != nil {
						log.Printf("Error echoing registration message to %s: %s\n", regMsg.Name, err)
						return
					}

					BroadcastChannel <- BroadcastMsg{
						Btype: PLAYERS,
					}
					BroadcastChannel <- BroadcastMsg{
						Btype:  FIRST_DATA,
						Client: c,
					}
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
						//TODO Here we validate the parameters of the msg
						SimulationChannel <- simulation.SimulatorMessage{
							Type:  simulation.MARK_CELL,
							X:     cmdMsg.X,
							Y:     cmdMsg.Y,
							Color: player.color,
						}
						log.Printf("Marking cell at (%d, %d) with color %32b", cmdMsg.X, cmdMsg.Y, player.color)

					case message.CommandType_PLACE_RLE:
						player := clients[c]
						//TODO Here we validate the parameters of the msg
						log.Println("Received RLE")
						SimulationChannel <- simulation.SimulatorMessage{
							Type:  simulation.PLACE_RLE,
							X:     cmdMsg.X,
							Y:     cmdMsg.Y,
							Color: player.color,
							Info:  cmdMsg.Text,
						}
					case message.CommandType_CLEAR_BOARD:
						SimulationChannel <- simulation.SimulatorMessage{
							Type: simulation.CLEAR_BOARD,
						}
					}
				}
			default:
				log.Printf("Received non-recognized message of type %d with content: %s", msg.Type, msg.Content)
			}
		}
	}
}
