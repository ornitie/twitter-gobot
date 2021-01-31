package server

import (
	"database/sql"
	"fmt"
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
}

func NewServer(envs map[string]string) (Server, error) {
	baseResource := resources.NewBaseResource(envs["bearer"])
	db := initializeDatabase(envs)
	defer db.Close()
	sqlStatement := "SELECT name from groups"
	row := db.QueryRow(sqlStatement)
	var title string
	_ = row.Scan(&title)
	log.Printf("%s", title)

	err := db.Ping()
	if err != nil {
		panic(err)
	}

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

func initializeDatabase(envs map[string]string) *sql.DB {
	databaseInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envs["dbHost"],
		envs["dbPort"],
		envs["dbUser"],
		envs["dbPassword"],
		envs["dbName"])

	db, err := sql.Open("postgres", databaseInfo)

	if err != nil {
		panic(err)
	}

	return db
}
