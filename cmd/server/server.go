package server

import (
	"github.com/gorilla/mux"
	"github.com/ornitie/twitter-gobot/cmd/controllers"
	"github.com/ornitie/twitter-gobot/pkg/resources"
	"log"
	"net/http"
)

type api struct {
	router          http.Handler
	rulesController *controllers.RulesController
}

type Server interface {
	Router() http.Handler
	SetRouter(*mux.Router)
}

func NewServer(bearer string) (Server, error) {
	baseResource := resources.NewBaseResource(bearer)

	api := &api{
		rulesController: controllers.NewRulesController(baseResource),
	}

	error := mapRoutes(api)

	if error != nil {
		log.Printf("failed, lul %v", error)
		return nil, error
	}

	return api, nil
}

func (api *api) Router() http.Handler {
	log.Print("Server up and running")

	return api.router
}

func (api *api) SetRouter(r *mux.Router) {
	api.router = r
}
