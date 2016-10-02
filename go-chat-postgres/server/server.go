package server

import (
	"log"
	"net/http"
	"os"

	"github.com/eaigner/jet"
)

type Server struct {
	db *jet.Db
}

func NewServer(db *jet.Db) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", use(basicAuth(s.template_handler)))
	mux.HandleFunc("/logout", s.logout_handler)

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
