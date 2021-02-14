package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/ornitie/twitter-gobot/cmd/controllers"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type api struct {
	router           http.Handler
	rulesController  *controllers.RulesController
	tweetsController *controllers.TweetsController
}

type Server interface {
	Router() http.Handler
	SetRouter(*mux.Router)
	StreamTweets()
}

func NewServer(envs map[string]string, db *sql.DB) (Server, error) {
	baseResource := resources.NewBaseResource(envs["bearer"])

	api := &api{
		rulesController:  controllers.NewRulesController(baseResource),
		tweetsController: controllers.NewTweetsController(baseResource),
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

func (api *api) StreamTweets() {
	api.tweetsController.StreamTweets()
}
