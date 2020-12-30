package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
	SetRouter(*mux.Router)
}

func NewServer() (Server, error) {
	api := &api{}

	error := mapRoutes(api)

	if error != nil {
		log.Printf("failed, lul %v", error)
		return nil, error
	}

	return api, nil
}

func (api *api) Router() http.Handler {
	return api.router
}

func (api *api) SetRouter(r *mux.Router) {
	api.router = r
}
