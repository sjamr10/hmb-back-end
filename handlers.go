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

type chat struct {
	ReceiverID string
	Msg        string
}

type data struct {
	ID       string
	Location string
	Chat     chat
}

func readMsg(conn *websocket.Conn) data {
	var data data
	// Read in a message
	err := conn.ReadJSON(&data)
	if err != nil {
		log.Println(err)
		return data
	}
	fmt.Println(data)
	return data
}

func sendMsg(data data, conn *websocket.Conn) {
	if err := conn.WriteJSON(data); err != nil {
		log.Println(err)
		return
	}
}

// Define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
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
