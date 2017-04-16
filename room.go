package main

import "github.com/gorilla/websocket"

/*
room {
name string
allCh chan message
sockets []*websocket.Conn
}

*/

type roomModel struct {
	name      string
	messageCh chan message
	sockets   []*websocket.Conn
}

func roomCtrl(s string) *roomModel {
	sockets := []*websocket.Conn{}
	msgCh := make(chan message)
	return &roomModel{
		name:      s,
		messageCh: msgCh,
		sockets:   sockets,
	}
}
