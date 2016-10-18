package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/simonleung8/bluemix-demo/go-chat-postgres/db"
	"github.com/simonleung8/bluemix-demo/go-chat-postgres/utils"
)

type template_data struct {
	Chats template.HTML
	User  string
}

func (s *Server) root_handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s.chat_handler(w, req)
	} else if req.Method == "GET" {
		s.template_handler(w, req)
	}
}

func (s *Server) template_handler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl")
	utils.Must(err, "Error parsing template")

	result, err := db.GetChats(s.db)
	utils.Must(err, "Error getting chat messages")

	cookie, err := req.Cookie("friend-circle-user")
	utils.Must(err, "Error getting cookie")

	d := template_data{
		Chats: template.HTML(build_chat_text(result)),
		User:  strings.TrimLeft(cookie.String(), "friend-circle-user="),
	}
	utils.Must(t.Execute(w, d), "Error executing template")
}

func (s *Server) chat_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["user_name"]) == 0 || len(r.Form["chat_msg"]) == 0 {
		log.Println("No username or chat message is sent")
	} else {
		utils.Must(db.SendChat(s.db, r.Form["user_name"][0], r.Form["chat_msg"][0]), "Error saving chat message")
	}

	http.Redirect(w, r, "/", 301)
}

func (s *Server) upload_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	utils.Must(r.ParseMultipartForm(32<<20), "Error parsing multipart form: ")

	file, header, err := r.FormFile("photo")
	utils.Must(err, "Error reading uploaded photo")
	defer file.Close()

	hostUrl := utils.HostImage(file, header)
	println("url", hostUrl)
	r.ParseForm()

	// talk to watson here
	watsonClassifyUrl := "https://gateway-a.watsonplatform.net/visual-recognition/api/v3/classify?api_key=ff3b8ffab1e2115fee4dcee3ab141a4bc0434aca&url=" + hostUrl + "&version=2016-05-19"
	req, err := http.NewRequest("GET", watsonClassifyUrl, nil)
	utils.Must(err, "Error creating request for watson classify")

	client := &http.Client{}
	res, err := client.Do(req)
	utils.Must(err, "Error posting request to watson classify")

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	utils.Must(err, "Error reading response body from image host")

	watsonResult := utils.Watson_images{}
	utils.Must(json.Unmarshal(body, &watsonResult), "Error unmarshaling watson service result")
	utils.Must(db.PostImage(s.db, r.Form["user_name"][0], hostUrl, watsonResult), "Error saving uploaded image")

	http.Redirect(w, r, "/", 301)
}

func (s *Server) logout_handler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/", 401)
}
