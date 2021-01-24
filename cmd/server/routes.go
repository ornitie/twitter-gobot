package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type handler func(*http.Request) (interface{}, error)

func mapRoutes(api *api) error {
	router := mux.NewRouter()

	api.SetRouter(router)

	router.HandleFunc("/rules", api.genericRouter(api.rulesController.GetRules)).Methods(http.MethodGet)
	router.HandleFunc("/rules", api.genericRouter(api.rulesController.CreateRule)).Methods(http.MethodPost)
	router.HandleFunc("/rules/{ID:[a-zA-Z0-9_]+}", api.genericRouter(api.rulesController.DeleteRule)).Methods(http.MethodDelete)

	return nil
}

func (api *api) genericRouter(function handler) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		response, error := function(request)

		writer.Header().Set("Content-Type", "application/json")

		if error != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		encoding_error := json.NewEncoder(writer).Encode(response)

		if encoding_error != nil {
			log.Printf("Error response %v", encoding_error)
		}
	}
}
