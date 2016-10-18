package main

import (
	"log"

	"github.com/simonleung8/bluemix-demo/go-chat-postgres/db"
	"github.com/simonleung8/bluemix-demo/go-chat-postgres/server"
)

func main() {
	client := db.NewClient()
	err := db.SeedDB(client)
	if err != nil {
		log.Print("Error seeding DB: " + err.Error())
	}

	s := server.NewServer(client)
	s.Start()
}
