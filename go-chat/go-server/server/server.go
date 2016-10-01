package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	couchdb "github.com/fjl/go-couchdb"
)

type template_data struct {
	Chats template.HTML
}

type messages struct {
	Rows []msg_data `json:"rows"`
}

type msg_data struct {
	Time string `json:"key"`
	Msg  string `json:"value"`
	Name string `json:"id"`
}

type Server struct {
	db *couchdb.DB
}

func NewServer(db *couchdb.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", use(basicAuth(s.template_handler)))

	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	port := os.Getenv("PORT")
	log.Println("Server is listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func must(err error, msg string) {
	if err != nil {
		log.Fatal(msg + ": " + err.Error())
	}
}

func build_chat_text(msgs messages) string {
	var str string
	for _, r := range msgs.Rows {
		chat_time, err := time.Parse("2006-01-02 15:04 MST", r.Time)
		if err != nil {
			log.Println("error parsing time ", err)
		}

		str = str + fmt.Sprintf("<b>%s</b><span style='font-size:0.7em; color:#aaa;'> (%s) </span>: %s<br>", r.Name, chat_time.Format("Mon 15:04"), r.Msg)
	}
	return str
}
