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
	router := mux.NewRouter()

	handler := &Server{
		router: router,
	}

	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	authorV1APi := v1.NewAuthorAPI()
	authorV1APi.Register(apiV1Router)

	documentV1Api := v1.NewDocumentApi()
	documentV1Api.Register(apiV1Router)

	return handler
}

func (handler *Server) Run() {
	http.Handle("/", handler.router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
