package main

import "github.com/gorilla/websocket"

type hub struct {
	rooms      map[string]*roomModel
	dumpCh     chan message
	unregister chan *websocket.Conn
}

func hubCtrl() hub {
	return hub{
		rooms:      make(map[string]*roomModel),
		dumpCh:     make(chan message),
		unregister: make(chan *websocket.Conn),
	}
}
