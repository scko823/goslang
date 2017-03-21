package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var mockMessages []string
var counter = 3

func init() {
	mockMessages = []string{"message1", "message2"}
}

func main() {
	http.HandleFunc("/messages", messages)
	http.ListenAndServe(":8080", nil)
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
