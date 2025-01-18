package main

import (
	"go-chat-app/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", websocket.HandleConnections)

	// Start the server
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
