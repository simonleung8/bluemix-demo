package server

import (
	"html/template"
	"log"
	"net/http"
)

func (s *Server) template_handler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl")
	must(err, "Error parsing template")

	var result messages
	err = s.db.View("_design/friends-circle", "get-msg", &result, nil)
	if err != nil {
		log.Print("error getting view", err.Error())
	}

	d := template_data{
		Chats: template.HTML(build_chat_text(result)),
	}
	must(t.Execute(w, d), "Error executing template")
}
