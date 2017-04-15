package main

import "github.com/gorilla/websocket"

/*
room {
name string
allCh chan message
sockets []*websocket.Conn
}

*/

type room struct {
	name      string
	messageCh chan message
	sockets   []*websocket.Conn
}

func roomCtrl(s string) room {
	var sockets []*websocket.Conn
	msgCh := make(chan message)
	return room{
		name:      s,
		messageCh: msgCh,
		sockets:   sockets,
	}
}
