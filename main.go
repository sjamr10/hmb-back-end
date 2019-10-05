package main

import (
	"fmt"
	"log"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ws", WsEndpoint)
}

func main() {
	fmt.Println("Starting server")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
