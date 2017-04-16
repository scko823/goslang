package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("There is an error attempting to promote to websocket")
		log.Printf("err is: %v\n", err)
		return
	}
	queries, _ := url.ParseQuery(r.URL.RawQuery)
	roomName := queries["room"][0]
	log.Printf("there is a new socket joining %v, \n", roomName)
	if room, ok := mainHub.rooms[roomName]; ok {
		newSockets := append(room.sockets, conn)
		room.sockets = newSockets
		fmt.Printf("len of rooms socket after: %v\n", len(room.sockets))
	} else {
		newRoom := roomCtrl(roomName)
		newSocketSlice := append(newRoom.sockets, conn)
		newRoom.sockets = newSocketSlice
		rooms[roomName] = newRoom
		fmt.Printf("newly created room: %v\n", roomName)
	}
	defer func(c *websocket.Conn) {
		rooms[roomName].unregister <- c
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
		rooms[roomName].messageCh <- *newMsg
		log.Println("messages.go line 47")
	}
}
