package main

import (
	"log"

	"github.com/simonleung8/bluemix-demo/go-chat/go-server/db"
	"github.com/simonleung8/bluemix-demo/go-chat/go-server/server"
)

func main() {
	cleint := db.NewClient()

	db, err := cleint.CreateDB("friendcircle")
	if err != nil {
		log.Print("Error creating/retriving DB: " + err.Error())
	}

	s := server.NewServer(db)
	s.Start()
}
