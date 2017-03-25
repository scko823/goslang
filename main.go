package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var mockMessages []string
var counter = 3
var tmpl *template.Template
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var sockets []*websocket.Conn
var dumpCh chan []byte

func init() {
	dumpCh = make(chan []byte)
	var err error
	mockMessages = []string{"message1", "message2"}
	tmpl, err = template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalf("there is an error when trying to parse templates. Err: %v\n", err)
	}

}

func main() {
	go func() {
		log.Print("running go routine to firehose all sockets")
		var byteMessage []byte
		for {
			byteMessage = <-dumpCh
			log.Println("Got new message")
			log.Println(string(byteMessage))
			for _, socket := range sockets {
				socket.WriteMessage(websocket.TextMessage, byteMessage)
			}
		}
	}()
	http.HandleFunc("/", index)
	http.HandleFunc("/messages", messages)
	http.HandleFunc("/ws", wsHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("an error occured when trying to start the server, %v\n", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", struct {
		Messages []string
	}{
		mockMessages,
	})
	return
}
