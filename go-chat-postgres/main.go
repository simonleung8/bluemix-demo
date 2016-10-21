package main

import (
	"log"
	"time"

	"github.com/simonleung8/bluemix-demo/go-chat-postgres/db"
	"github.com/simonleung8/bluemix-demo/go-chat-postgres/server"
)

func main() {
	//sleep a little for logs to start streaming
	time.Sleep(5 * time.Second)
	databaseClient, err := db.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	err = db.SeedDB(databaseClient)
	if err != nil {
		log.Fatal("Error seeding DB: " + err.Error())
	}

	s := server.NewServer(databaseClient)
	s.Start()
}
