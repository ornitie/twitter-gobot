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
	response, error := resources.Get("https://pokeapi.co/api/v2/ability/?limit=20&offset=20")
	if error != nil {
		fmt.Println("Error Calling resource")
		return
	}

	pokeresponse := &PokeResponse{}
	_ = resources.ToStruct(*response, pokeresponse)
	fmt.Println(pokeresponse.Count)
}
