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
	router.HandleFunc("/author/hello", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello world")
	})
}
