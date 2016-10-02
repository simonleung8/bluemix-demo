package db

import (
	"time"

	"github.com/eaigner/jet"
)

type Chats []struct {
	Id    int
	Name  string    `json:"name"`
	Chat  string    `json:"chat"`
	Added time.Time `json:"added"`
}

func GetChats(client *jet.Db) (Chats, error) {
	var result Chats
	err := client.Query("SELECT * FROM chats ORDER BY added").Rows(&result)
	return result, err
}
