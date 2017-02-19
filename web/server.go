package web

import (
	"github.com/geryheselmans/go-test-server/model"
	"github.com/geryheselmans/go-test-server/web/api/v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router           *mux.Router
	authorRepository models.AuthorRepository
}

func New(authorRepository models.AuthorRepository) *Server {
	router := mux.NewRouter()

	server := &Server{
		router:           router,
		authorRepository: authorRepository,
	}

	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	authorV1APi := v1.NewAuthorAPI(authorRepository)
	authorV1APi.Register(apiV1Router)

	documentV1Api := v1.NewDocumentApi()
	documentV1Api.Register(apiV1Router)

	return server
}

func (handler *Server) Run() {
	http.Handle("/", handler.router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
