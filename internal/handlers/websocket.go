package handlers

import (
	"distributed-systems-chatbot/internal/gpt"
	"distributed-systems-chatbot/internal/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("WebSocket upgrade failed:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			panic(err)
		}
	}(ws)

	clients[ws] = true

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		response := gpt.GetResponse(msg.Message)
		gptMsg := models.Message{
			Username: "Viegas AI",
			Message:  response,
		}
		for client := range clients {
			err := client.WriteJSON(gptMsg)
			if err != nil {
				log.Printf("error: %v", err)
				err := client.Close()
				if err != nil {
					return
				}
				delete(clients, client)
			}
		}
	}
}
