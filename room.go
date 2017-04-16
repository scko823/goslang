package main

import (
	"log"

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
