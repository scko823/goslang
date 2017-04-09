package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("There is an error attempting to promote to websocket")
		log.Printf("err is: %v\n", err)
		return
	}
	sockets = append(sockets, conn)
	defer func(c *websocket.Conn) {
		unregister <- c
		c.Close()
	}(conn)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		newMsg := new(message)
		json.Unmarshal(p, newMsg)
		dumpCh <- *newMsg
		log.Println("messages.go line 28")
	}
}
