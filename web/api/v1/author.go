package v1

import (
	"encoding/json"
	"github.com/geryheselmans/go-test-server/model"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthorAPI struct {
	authorRepository models.AuthorRepository
}

func NewAuthorAPI(authorRepository models.AuthorRepository) *AuthorAPI {
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
	authors, err1 := api.authorRepository.FindAll()

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(response)

	err2 := encoder.Encode(authors)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (api *AuthorAPI) GetAuthorById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	author, err1 := api.authorRepository.FindByAuthorName(vars["authorName"])

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if author == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(response)

	err2 := encoder.Encode(author)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (api *AuthorAPI) CreateAuthor(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var author models.Author
	err1 := decoder.Decode(&author)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err2 := api.authorRepository.Save(author.AuthorName, &author)
	if err2 != nil {
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
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

func (api *AuthorAPI) DeleteAuthorById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	err := api.authorRepository.Delete(vars["authorName"])

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
