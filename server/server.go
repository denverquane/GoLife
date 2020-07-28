package main

import (
	"flag"
	"github.com/denverquane/golife/proto/message"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

var addr = flag.String("addr", ":5000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	Run(addr)
}

func Run(addr *string) {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
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
