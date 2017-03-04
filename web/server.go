package web

import (
	"github.com/geryheselmans/go-test-server/model"
	"github.com/geryheselmans/go-test-server/web/api/v1"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	router           *mux.Router
	authorRepository model.AuthorRepository
	errChan          chan error
}

var log *zap.Logger

func New(logger *zap.Logger, authorRepository model.AuthorRepository) *Server {
	log = logger

	router := mux.NewRouter()

	server := &Server{
		router:           router,
		authorRepository: authorRepository,
		errChan:          make(chan error),
	}

	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	authorV1APi := v1.NewAuthorAPI(log, authorRepository)
	authorV1APi.Register(apiV1Router)

	return server
}

func (handler *Server) Run() {
	loggingMiddleware := loggingMiddleware(handler.router)
	http.Handle("/", loggingMiddleware)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		handler.errChan <- err
	}
}

func (handler *Server) ErrChan() chan error {
	return handler.errChan
}
