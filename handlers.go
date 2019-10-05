package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

var connections = make(map[string]*websocket.Conn)

type msg struct {
	Type string
	Data string
}

func readMsg(conn *websocket.Conn) msg {
	var msg msg
	// Read in a message
	err := conn.ReadJSON(&msg)
	if err != nil {
		log.Println(err)
		return msg
	}
	fmt.Println(msg)
	return msg
}

func sendMsg(msg msg, conn *websocket.Conn) {
	if err := conn.WriteJSON(msg); err != nil {
		log.Println(err)
		return
	}
}

// Define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	connections["user1"] = conn

	msg := msg{Type: "id", Data: "user1"}
	sendMsg(msg, conn)
	for {
		data := readMsg(conn)
		sendMsg(data, conn)
	}
}

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WsEndpoint ...
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// Upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	// Listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}
