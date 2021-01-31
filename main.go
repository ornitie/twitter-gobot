package main

import (
	// "bufio"
	// "fmt"
	// "github.com/joho/godotenv"
	// "github.com/ornitie/twitter-gobot/internal/models"

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

		s, error := server.NewServer(envs)

		if error != nil {
			return
		}

		error = http.ListenAndServe(":8080", s.Router())

		if error != nil {
			log.Fatalf("Error creating the server %v", error)
		}
	}

	fmt.Println("Successfully connected!")

	log.Fatalf("Error failed missing env var")
	// fmt.Println("hello world, This is my own twitter-gobot")
	// err := godotenv.Load()
	// value, defined := os.LookupEnv("SOMETHING_ELSE")
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	// if defined {
	// 	fmt.Println(value)
	// }
	// createRule := models.CreateRule{Add: []models.Rule{models.Rule{Value: "from:shoe0nhead"}}}
	// 	post_response, _ := baseResource.Post("https://api.twitter.com/2/tweets/search/stream/rules", createRule)
	// 	fmt.Printf("GOT ONE %+v", post_response)
	// response, _ := baseResource.Get("https://api.twitter.com/2/tweets/search/stream")
	// fmt.Println("Response: Content-length:", response.Header.Get("Content-length"))

	// reader := bufio.NewReader(response.Body)
	// for {
	// 	// fmt.Printf("YEA %v", response.Body)
	// 	line, err := reader.ReadBytes('\n')
	// 	if err != nil {
	// 		fmt.Print("FUCK")
	// 	}
	// 	fmt.Printf("GOT ONE %v", string(line))
	// }
	// if err != nil {
	// 	fmt.Println("Error Calling resource")
	// 	return
	// }

	// pokeresponse := &PokeResponse{}
	// _ = resources.ToStruct(*response, pokeresponse)
	// fmt.Println(pokeresponse.Data.Username)
	// }
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
