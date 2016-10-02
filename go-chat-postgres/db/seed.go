package db

import (
	"fmt"
	"log"

	"github.com/eaigner/jet"
)

func SeedDB(client *jet.Db) error {
	// drop_table("chats", client)
	return seed_chats_table(client)
}

func drop_table(tableName string, client *jet.Db) {
	err := client.Query(fmt.Sprintf("drop table %s;", tableName)).Run()
	if err != nil {
		log.Print(fmt.Sprintf("Error dropping table %s: %s", tableName, err.Error()))
	}
}

func seed_chats_table(client *jet.Db) error {
	return client.Query(`
		CREATE TABLE IF NOT EXISTS "chats" (
			"id" SERIAL PRIMARY KEY,
			"name" varchar(15) NOT NULL,
			"chat" text NOT NULL,
			"added" timestamp default (now() at time zone 'utc')
		);`).Run()
}
