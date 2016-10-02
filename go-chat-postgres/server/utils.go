package server

import (
	"fmt"

	"github.com/simonleung8/bluemix-demo/go-chat-postgres/db"
)

func build_chat_text(chats db.Chats) string {
	var str string
	for _, r := range chats {
		str = str + fmt.Sprintf("<b>%s</b><span style='font-size:0.7em; color:#aaa;'> (%s utc) </span>: %s<br>", r.Name, r.Added.Format("Mon 15:04"), r.Chat)
	}
	return str
}
