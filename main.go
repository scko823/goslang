package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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
var rooms map[string]*roomModel
var mainHub hub

func init() {
	// makeRedisConn()
	mainHub = hubCtrl()
	dumpCh = mainHub.dumpCh
	unregister = mainHub.unregister
	rooms = mainHub.rooms
	rooms["main"] = roomCtrl("main")
}

func main() {
	go func() {
		log.Print("running go routine to firehose all sockets")
		// var newMsg message
		for {
			select {
			case _ = <-dumpCh:

			}
		}
	}()
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html")
	})
	http.HandleFunc("/ws", wsHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("an error occured when trying to start the server, %v\n", err)
	}
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "index.html", nil)
// 	return
// }
