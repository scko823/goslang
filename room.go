package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
room {
name string
allCh chan message
sockets []*websocket.Conn
}

*/

type roomModel struct {
	name       string
	messageCh  chan message
	sockets    []*websocket.Conn
	unregister chan *websocket.Conn
}

func roomCtrl(s string) *roomModel {
	sockets := []*websocket.Conn{}
	msgCh := make(chan message)
	unregister := make(chan *websocket.Conn)
	newRoom := roomModel{
		name:       s,
		messageCh:  msgCh,
		sockets:    sockets,
		unregister: unregister,
	}
	go func() {
		for {
			select {
			case newMsg := <-msgCh:
				log.Printf("message is for room:%s\n", s)
				log.Printf("Len of sockets: %v \n", len(newRoom.sockets))
				for _, socket := range newRoom.sockets {
					socket.WriteJSON(newMsg)
				}
			case leftConn := <-unregister:
				for i, socket := range newRoom.sockets {
					if socket == leftConn {
						sockets := append(newRoom.sockets[:i], newRoom.sockets[i+1:]...)
						log.Printf("a socket is leaving...\nnew len of sockets: %v", len(sockets))
						newRoom.sockets = sockets
						break
					}
				}
			}

		}
	}()
	return &newRoom
}

type roomEvent struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func roomListener(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("There is an error attempting to promote to websocket on /rooms-channel")
		log.Printf("err is: %v\n", err)
		return
	}
	allSockets = append(allSockets, conn)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("There is an error attempting to read message on websocket on /rooms-channel")
			log.Printf("err is: %v\n", err)
			break
		}
		fmt.Println("get new roomEvent")
		fmt.Printf("%s\n", string(p))
		newRoomEvent := new(roomEvent)
		json.Unmarshal(p, newRoomEvent)
		roomChannel <- *newRoomEvent
	}
}
