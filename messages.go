package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

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
		log.Println(string(p))
		dumpCh <- p
		log.Println("messages.go line 28")
	}
}

func messages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		marshalledMsg, err := json.Marshal(mockMessages)
		if err != nil {
			log.Println("an error occur during JSON marshal process")
			log.Println(err)
		}
		io.WriteString(w, string(marshalledMsg))
		return
	} else if r.Method == http.MethodPost {
		newMsgByte := []byte{}
		_, err := r.Body.Read(newMsgByte)
		if err != nil {
			log.Println("an error occur when reading request Body")
			log.Printf("The request body is\n")
			log.Printf("%v", r.Body)
			log.Printf("The error is\n")
			log.Println(err)
		}
		mockMessages = append(mockMessages, "message"+strconv.Itoa(counter))
		counter++
		marshalledMsg, _ := json.Marshal(mockMessages)
		if err != nil {
			counter--
			log.Println("an error occur during JSON marshal process")
			log.Println(err)
		}
		io.WriteString(w, string(marshalledMsg))
		return
	}
}
