package db

import (
	"fmt"
	"time"

	"github.com/eaigner/jet"
	"github.com/simonleung8/bluemix-demo/go-chat-postgres/utils"
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
	byteStr := utils.BuildByteString(chat)
	return client.Query(fmt.Sprintf("INSERT INTO chats (name, chat) VALUES ('%s', '%s');", user, byteStr)).Run()
}

func PostImage(client *jet.Db, user string, imgUrl string) error {
	msg := fmt.Sprintf(`<img src="%s" width="200px" align="middle">`, imgUrl)
	return client.Query(fmt.Sprintf(`INSERT INTO chats (name, chat) VALUES ('%s', '%s');`, user, msg)).Run()
}
