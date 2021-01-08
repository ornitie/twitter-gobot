package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func mapRoutes(a *api) error {
	router := mux.NewRouter()

	a.SetRouter(router)

	router.HandleFunc("/rules", a.getRules).Methods(http.MethodGet)

	return nil
}

func (api *api) getRules(writer http.ResponseWriter, request *http.Request) {
	// gophers := map[string]string{
	// 	"response": "nope, chuck testa in the routes",
	// }

	gophers := api.rulesController.GetRules()

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(gophers)
}
