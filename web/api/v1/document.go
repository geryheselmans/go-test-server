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
	router.HandleFunc("/document/hello", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello world")
	})
}
