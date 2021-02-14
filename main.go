package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ornitie/twitter-gobot/cmd/server"
)

func main() {
	envs, defined := loadEnvs()

	if defined {

		db := initializeDatabase(envs)
		defer db.Close()

		err := db.Ping()
		if err != nil {
			panic(err)
		}

		s, error := server.NewServer(envs)

		if error != nil {
			return
		}

		go func() {
			s.StreamTweets()
		}()

		error = http.ListenAndServe(":8080", s.Router())

		if error != nil {
			log.Fatalf("Error creating the server %v", error)
		}
	}

	log.Fatalf("Error failed missing env var")
}

func loadEnvs() (map[string]string, bool) {
	_ = godotenv.Load()

	envs := map[string]string{}

	var defined bool

	if envs["bearer"], defined = os.LookupEnv("BEARER_TOKEN"); !defined {
		return envs, false
	}
	if envs["dbHost"], defined = os.LookupEnv("DATABASE_HOST"); !defined {
		return envs, false
	}
	if envs["dbPort"], defined = os.LookupEnv("DATABASE_PORT"); !defined {
		return envs, false
	}
	if envs["dbUser"], defined = os.LookupEnv("DATABASE_USER"); !defined {
		return envs, false
	}
	if envs["dbPassword"], defined = os.LookupEnv("DATABASE_PASSWORD"); !defined {
		return envs, false
	}
	if envs["dbName"], defined = os.LookupEnv("DATABASE_NAME"); !defined {
		return envs, false
	}

	return envs, true
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
