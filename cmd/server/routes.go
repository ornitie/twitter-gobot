package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(*http.Request) interface{}

func mapRoutes(api *api) error {
	router := mux.NewRouter()

	api.SetRouter(router)

	router.HandleFunc("/rules", api.genericRouter(api.rulesController.GetRules)).Methods(http.MethodGet)

	return nil
}

func (api *api) genericRouter(function handler) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, reader *http.Request) {
		rules := function(reader)

		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(rules)
	}
}
