package server

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/simonleung8/bluemix-demo/go-chat-postgres/db"
)

type template_data struct {
	Chats template.HTML
	User  string
}

func (s *Server) template_handler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl")
	must(err, "Error parsing template")

	result, err := db.GetChats(s.db)
	if err != nil {
		log.Print("error getting chats: ", err.Error())
	}

	cookie, err := req.Cookie("friend-circle-user")
	if err != nil {
		log.Print("error getting cookie", err.Error())
	}

	d := template_data{
		Chats: template.HTML(build_chat_text(result)),
		User:  strings.TrimLeft(cookie.String(), "friend-circle-user="),
	}
	must(t.Execute(w, d), "Error executing template")
}

func (s *Server) logout_handler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/", 401)
}
