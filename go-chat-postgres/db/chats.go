package db

import (
	"fmt"
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

func SendChat(client *jet.Db, user string, chat string) error {
	return client.Query(fmt.Sprintf(`INSERT INTO chats (name, chat) VALUES ('%s', '%s');`, user, chat)).Run()
}

func PostImage(client *jet.Db, user string, imgUrl string) error {
	msg := fmt.Sprintf(`<img src="%s" width="200px" align="middle">`, imgUrl)
	return client.Query(fmt.Sprintf(`INSERT INTO chats (name, chat) VALUES ('%s', '%s');`, user, msg)).Run()
}
