package web

import (
	"github.com/geryheselmans/go-test-server/model"
	"github.com/geryheselmans/go-test-server/web/api/v1"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router           *mux.Router
	authorRepository model.AuthorRepository
}

func New(authorRepository model.AuthorRepository) *Server {
	router := mux.NewRouter()

	server := &Server{
		router:           router,
		authorRepository: authorRepository,
	}

	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	authorV1APi := v1.NewAuthorAPI(authorRepository)
	authorV1APi.Register(apiV1Router)

	return server
}

func (handler *Server) Run() {
	http.Handle("/", handler.router)

	err := http.ListenAndServe(":8080", nil)
	log.WithError(err).Fatal("Error while starting serivce")
}
