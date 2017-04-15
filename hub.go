package main

import "github.com/gorilla/websocket"

type hub struct {
	rooms      map[string]room
	dumpCh     chan message
	unregister chan *websocket.Conn
}

func hubCtrl() hub {
	return hub{
		rooms:      make(map[string]room),
		dumpCh:     make(chan message),
		unregister: make(chan *websocket.Conn),
	}
}
