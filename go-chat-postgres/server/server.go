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

	mux.HandleFunc("/", use(basicAuth(s.root_handler))) //handles both GET and POST
	mux.HandleFunc("/upload", s.upload_handler)         //handles POST
	mux.HandleFunc("/chats", s.get_chats_handler)       //handles GET
	mux.HandleFunc("/logout", s.logout_handler)

	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
