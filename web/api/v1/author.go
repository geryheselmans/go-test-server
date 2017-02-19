package v1

import (
	"fmt"
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
	router.HandleFunc("/authors", GetAllAuthors).
		Methods("GET").
		HeadersRegexp("Accept", "application/(xml|json)")

	router.HandleFunc("/authors/{authorName:[a-z0-9]+}", GetAuthorById).
		Methods("GET").
		HeadersRegexp("Accept", "application/(xml|json)")

	router.HandleFunc("/authors", CreateAuthor).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/(xml|json)")

	router.HandleFunc("/authors/{authorName:[a-z0-9]+}", UpdateAuthorById).
		Methods("PUT").
		HeadersRegexp("Content-Type", "application/(xml|json)")

	router.HandleFunc("/authors", DeleteAllAuthors).
		Methods("DELETE")

	router.HandleFunc("/authors/{authorName:[a-z0-9]+}", DeleteAuthorById).
		Methods("DELETE")
}

func GetAllAuthors(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Header().Add("Content-Type", "application/json")

	fmt.Fprint(response, "{\"test\":\"Hello world\"}")
}

func GetAuthorById(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(request)

	fmt.Fprintf(response, "{\"test\":\"%s\"}", vars["authorName"])
}

func CreateAuthor(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusCreated)
}

func UpdateAuthorById(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func DeleteAllAuthors(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func DeleteAuthorById(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}
