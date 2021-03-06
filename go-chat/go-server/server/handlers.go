package server

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type template_data struct {
	Chats template.HTML
	User  string
}

type test_data struct {
	Chats []test2 `json:"chats"`
}
type test2 struct {
	Chat string `json:"msg"`
	Time string `json:"time"`
}

func (s *Server) template_handler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl")
	must(err, "Error parsing template")

	dd := test_data{
		Chats: []test2{
			test2{
				Chat: "msg...123",
				Time: "",
			},
		},
	}
	aa, ee := s.db.Put("test", dd, "5-baec9522c4f7067d352c6e08c06de838")
	must(ee, "Error blah blah template")
	println(aa)

	var result messages
	err = s.db.View("_design/friends-circle", "get-msg", &result, nil)
	if err != nil {
		log.Print("error getting view", err.Error())
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
