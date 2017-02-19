package web

import (
	"github.com/geryheselmans/go-test-server/web/api/v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func New() *Server {
	r := mux.NewRouter()

	handler := &Server{
		router: r,
	}

	authorAPi := v1.NewAuthorAPI()
	authorAPi.Register(r.PathPrefix("/api/v1").Subrouter())

	return handler
}

func (handler *Server) Run() {
	http.Handle("/", handler.router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
