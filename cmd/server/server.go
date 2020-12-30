package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
	api.router = router

	router.HandleFunc("/rules", api.getRules).Methods(http.MethodGet)

	return api
}

func (api *api) Router() http.Handler {
	return api.router
}

func (api *api) getRules(writer http.ResponseWriter, request *http.Request) {
	gophers := map[string]string{
		"response": "nope, chuck testa",
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(gophers)
}
