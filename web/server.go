package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

type Server struct {
	router *mux.Router
}

func New() *Server {
	r := mux.NewRouter()

	handler := &Server{
		router: r,
	}

	r.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello world")
	})

	return handler
}

func (handler *Server) Run() {
	http.Handle("/", handler.router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
