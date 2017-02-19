package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthorAPI struct {
}

func NewAuthorAPI() *AuthorAPI {
	return &AuthorAPI{}
}

func (api *AuthorAPI) Register(router *mux.Router) {
	router.HandleFunc("/authors", GetAllAuthors).
		Methods("GET").
		HeadersRegexp("Accept", "application/(xml|json)")

	router.HandleFunc("/authors/{authorId:[0-9]+}", GetAuthorById).
		Methods("GET").
		HeadersRegexp("Accept", "application/(xml|json)")

	router.HandleFunc("/authors", CreateAuthor).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/(xml|json)")

	router.HandleFunc("/authors/{authorId:[0-9]+}", UpdateAuthorById).
		Methods("PUT").
		HeadersRegexp("Content-Type", "application/(xml|json)")

	router.HandleFunc("/authors", DeleteAllAuthors).
		Methods("DELETE")

	router.HandleFunc("/authors/{authorId:[0-9]+}", DeleteAuthorById).
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

	fmt.Fprintf(response, "{\"test\":\"%s\"}", vars["authorId"])
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
