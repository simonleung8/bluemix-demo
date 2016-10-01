package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	couchdb "github.com/fjl/go-couchdb"
)

type messages struct {
	Rows []msg_data `json:"rows"`
}

type msg_data struct {
	Time  string `json:"key"`
	Chats string `json:"value"`
	Name  string `json:"id"`
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
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("templates/index.tmpl")
		must(err, "Error parsing template")

		// var result msg_data
		var result messages
		err = s.db.View("_design/friends-circle", "get-msg", &result, nil)
		d := msg_data{
			// Chats: `Testing
			// hihi
			// `,
			Chats: fmt.Sprintf("%#v", result),
		}
		must(t.Execute(w, d), "Error executing template")
	})

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
