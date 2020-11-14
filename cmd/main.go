package main

import (
	"encoding/json"
	"fmt"
	"github.com/ornitie/twitter-gobot/pkg"
)

type (
	PokeResponse struct {
		Count int64 `json:"count" bson:"count"`
	}
)

func main() {
	fmt.Println("hello world, This is my own twitter-gobot")
	pokeresponse := &PokeResponse{}
	resp, err := resources.Get("https://pokeapi.co/api/v2/ability/?limit=20&offset=20")
	if err != nil {
		fmt.Println(err)
	} else {
		err := json.NewDecoder(resp.Body).Decode(pokeresponse)
		if err == nil {
			fmt.Println(pokeresponse.Count)
		}
	}
}
