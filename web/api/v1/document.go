package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type DocumentApi struct {
}

func NewDocumentApi() *DocumentApi {
	return &DocumentApi{}
}

func (api *DocumentApi) Register(router *mux.Router) {
	router.HandleFunc("/authors/{authorId:[0-9]+}/documents", GetAllDocumentsByAuthor).
		Methods("GET").
		HeadersRegexp("Accept", "application/(xml|json)")

	router.HandleFunc("/authors/{authorId:[0-9]+}/documents/{documentId:[0-9]+}", GetDocumentByIdAndAuthor).
		Methods("GET").
		HeadersRegexp("Accept", "application/(xml|json)")

	router.HandleFunc("/authors/{authorId:[0-9]+}/documents", CreateDocumentByAuthor).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/(xml|json)")

	router.HandleFunc("/authors/{authorId:[0-9]+}/documents/{documentId:[0-9]+}", UpdateDocumentByAuthor).
		Methods("PUT").
		HeadersRegexp("Content-Type", "application/(xml|json)")

	router.HandleFunc("/authors/{authorId:[0-9]+}/documents", DeleteAllDocumentsByAuthor).
		Methods("DELETE")

	router.HandleFunc("/authors/{authorId:[0-9]+}/documents/{id:[0-9]+}", DeleteDocumentByAuthorAndId).
		Methods("DELETE")
}

func GetAllDocumentsByAuthor(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Header().Add("Content-Type", "application/json")

	fmt.Fprint(response, "{\"test\":\"Hello world\"}")
}

func GetDocumentByIdAndAuthor(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(request)

	fmt.Fprintf(response, "{\"test\":\"%s\"}", vars["documentId"])
}

func UpdateDocumentByAuthor(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusCreated)
}

func CreateDocumentByAuthor(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func DeleteDocumentByAuthorAndId(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func DeleteAllDocumentsByAuthor(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}
