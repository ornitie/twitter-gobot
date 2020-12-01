package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ornitie/twitter-gobot/pkg"
	"os"
)

type (
	PokeResponse struct {
		Count int64 `json:"count" bson:"count"`
		Data  struct {
			Username string `json:"username" bson:"username"`
		} `json:"data" bson:"data"`
	}
)

func main() {
	fmt.Println("hello world, This is my own twitter-gobot")
	err := godotenv.Load()
	value, defined := os.LookupEnv("SOMETHING_ELSE")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	if defined {
		fmt.Println(value)
	}
	bearer, defined := os.LookupEnv("BEARER_TOKEN")
	if defined {
		baseResource := resources.NewBaseResource(bearer)
		response, error := baseResource.Get("https://api.twitter.com/2/users/by/username/negromano")
		if error != nil {
			fmt.Println("Error Calling resource")
			return
		}

		pokeresponse := &PokeResponse{}
		_ = resources.ToStruct(*response, pokeresponse)
		fmt.Println(pokeresponse.Data.Username)
	}
}
