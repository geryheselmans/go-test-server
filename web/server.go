package web

import (
	"github.com/geryheselmans/go-test-server/model"
	"github.com/geryheselmans/go-test-server/web/api/v1"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router           *mux.Router
	authorRepository model.AuthorRepository
	errChan          chan error
}

func New(authorRepository model.AuthorRepository) *Server {
	router := mux.NewRouter()

	server := &Server{
		router:           router,
		authorRepository: authorRepository,
		errChan:          make(chan error),
	}

	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	authorV1APi := v1.NewAuthorAPI(authorRepository)
	authorV1APi.Register(apiV1Router)

	return server
}

func (handler *Server) Run() {
	http.Handle("/", handler.router)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		handler.errChan <- err
	}
}

func (handler *Server) ErrChan() chan error {
	return handler.errChan
}
