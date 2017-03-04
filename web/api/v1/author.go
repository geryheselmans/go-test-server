package v1

import (
	"encoding/json"
	"github.com/geryheselmans/go-test-server/model"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type AuthorAPI struct {
	authorRepository model.AuthorRepository
}

var log *zap.Logger

func NewAuthorAPI(log *zap.Logger, authorRepository model.AuthorRepository) *AuthorAPI {
	log = log

	return &AuthorAPI{
		authorRepository: authorRepository,
	}
}

func (api *AuthorAPI) Register(router *mux.Router) {
	router.HandleFunc("/authors", api.GetAllAuthors).
		Methods("GET").
		Headers("Accept", "application/json")

	router.HandleFunc("/authors/{authorName:[a-z0-9]+}", api.GetAuthorById).
		Methods("GET").
		Headers("Accept", "application/json")

	router.HandleFunc("/authors", api.CreateAuthor).
		Methods("POST").
		Headers("Content-Type", "application/json")

	router.HandleFunc("/authors/{authorName:[a-z0-9]+}", api.UpdateAuthorById).
		Methods("PUT").
		Headers("Content-Type", "application/json")

	router.HandleFunc("/authors", api.DeleteAllAuthors).
		Methods("DELETE")

	router.HandleFunc("/authors/{authorName:[a-z0-9]+}", api.DeleteAuthorById).
		Methods("DELETE")
}

func (api *AuthorAPI) GetAllAuthors(response http.ResponseWriter, request *http.Request) {
	authors, err := api.authorRepository.FindAll()

	if err != nil {
		log.Error("Error with find all authors", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)

		return
	}

	encoder := json.NewEncoder(response)

	err = encoder.Encode(authors)
	if err != nil {
		log.Error("Error with json encode", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (api *AuthorAPI) GetAuthorById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	author, err := api.authorRepository.FindByAuthorName(vars["authorName"])

	if err != nil {
		log.Error("Error with find by author name", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)

		return
	}

	if author == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(response)

	err = encoder.Encode(author)
	if err != nil {
		log.Error("Error with json encode", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (api *AuthorAPI) CreateAuthor(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var author model.Author
	err := decoder.Decode(&author)

	if err != nil {
		log.Error("Error with json decode", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = author.Save(api.authorRepository)
	if err != nil {
		log.Error("Error with save author", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)

		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (api *AuthorAPI) UpdateAuthorById(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func (api *AuthorAPI) DeleteAllAuthors(response http.ResponseWriter, request *http.Request) {
	err := api.authorRepository.Clear()

	if err != nil {
		log.Error("Error with delete all", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

func (api *AuthorAPI) DeleteAuthorById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	err := api.authorRepository.Delete(vars["authorName"])

	if err != nil {
		log.Error("Error with delete author", zap.Error(err))

		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
