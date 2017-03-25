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

func init() {
	var err error
	mockMessages = []string{"message1", "message2"}
	tmpl, err = template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalf("there is an error when trying to parse templates. Err: %v\n", err)
	}
}

func main() {
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
