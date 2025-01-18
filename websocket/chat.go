package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Register the new client
	clients[conn] = true
	log.Println("New client connected")

	for {
		// Read message from the client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}

		// Broadcast the message to all clients
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
