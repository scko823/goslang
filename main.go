package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var tmpl *template.Template
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var sockets []*websocket.Conn
var unregister chan *websocket.Conn

type message struct {
	MessageText string `json:"message"`
	Room        string `json:"room"`
	Timestamp   int64  `json:"time"`
}

var dumpCh chan message
var messages = []message{}
var rooms map[string]room

func init() {
	mainHub := hubCtrl()
	dumpCh = mainHub.dumpCh
	unregister = mainHub.unregister
	rooms := mainHub.rooms
	rooms["main"] = roomCtrl("main")
	var err error
	tmpl, err = template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalf("there is an error when trying to parse templates. Err: %v\n", err)
	}
}

func main() {
	go func() {
		log.Print("running go routine to firehose all sockets")
		var newMsg message
		for {
			select {
			case newMsg = <-dumpCh:
				log.Println("Got new message:")
				log.Printf("message is for room: %s\n", newMsg.Room)
				log.Printf("Len of sockets: %v \n", len(sockets))
				for _, socket := range sockets {
					socket.WriteJSON(newMsg)
				}
			case leftConn := <-unregister:
				for i, socket := range sockets {
					if socket == leftConn {
						sockets = append(sockets[:i], sockets[i+1:]...)
						log.Printf("a socket is leaving...\nnew len of sockets: %v", len(sockets))
						break
					}
				}
			}

		}
	}()
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/ws", wsHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("an error occured when trying to start the server, %v\n", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	return
}
