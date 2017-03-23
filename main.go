package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

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
		// log.Fatalf("there is an error when trying to parse templates. Err: %v\n", err)
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

func wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("There is an error attempting to promote to websocket")
		fmt.Printf("err is: %v\n", err)
		return
	}
	for {
		var p string
		var err error
		err = conn.ReadJSON(&p)
		if err != nil {
			return
		}
		err = conn.WriteJSON(p)
		if err != nil {
			return
		}
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

func messages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		marshalledMsg, err := json.Marshal(mockMessages)
		if err != nil {
			fmt.Println("an error occur during JSON marshal process")
			fmt.Println(err)
		}
		io.WriteString(w, string(marshalledMsg))
		return
	} else if r.Method == http.MethodPost {
		newMsgByte := []byte{}
		_, err := r.Body.Read(newMsgByte)
		if err != nil {
			fmt.Println("an error occur when reading request Body")
			fmt.Printf("The request body is\n")
			fmt.Printf("%v", r.Body)
			fmt.Printf("The error is\n")
			fmt.Println(err)
		}
		mockMessages = append(mockMessages, "message"+strconv.Itoa(counter))
		counter++
		marshalledMsg, _ := json.Marshal(mockMessages)
		if err != nil {
			counter--
			fmt.Println("an error occur during JSON marshal process")
			fmt.Println(err)
		}
		io.WriteString(w, string(marshalledMsg))
		return
	}
}
