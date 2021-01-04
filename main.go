package main

import (
	// "bufio"
	// "fmt"
	// "github.com/joho/godotenv"
	// "github.com/ornitie/twitter-gobot/internal/models"
	"github.com/ornitie/twitter-gobot/cmd/server"
	"net/http"
	// "os"
)

func main() {
	s := server.NewServer()
	http.ListenAndServe(":8080", s.Router())
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
	// fmt.Printf("GOT ONE %+v", createRule)
	// bearer, defined := os.LookupEnv("BEARER_TOKEN")
	// if defined {
	// 	baseResource := resources.NewBaseResource(bearer)
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
