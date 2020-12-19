package routes

import (
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

	return api
}

func (api *api) Router() http.Handler {
	return api.router
}
