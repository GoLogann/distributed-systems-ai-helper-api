package main

import (
	"distributed-systems-chatbot/internal/config"
	"distributed-systems-chatbot/internal/handlers"
	"distributed-systems-chatbot/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	go handlers.HandleMessages()

	log.Println("Server starting on:", cfg.ServerAddress)
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed:", err)
	}
}
