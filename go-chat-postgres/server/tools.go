package server

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonleung8/bluemix-demo/go-chat-postgres/db"
)

func build_chat_text(chats db.Chats) string {
	var str string
	for _, r := range chats {

		chatStr := strings.Split(r.Chat, " ")
		var b []byte
		for _, s := range chatStr {
			i, _ := strconv.ParseInt(s, 10, 64)
			b = append(b, byte(i))
		}

		str = fmt.Sprintf(`
		<div>
			<b>%s</b><span style='font-size:0.7em; color:#aaa;'> (%s utc) </span>: %s<br>
		</div>`, r.Name, r.Added.Format("Mon 15:04"), string(b)) + str
	}
	return str
}
