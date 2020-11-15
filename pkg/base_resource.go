package resources

import (
	"encoding/json"
	"log"
	"net/http"
)

func Get(path string) (*http.Response, error) {
	log.Printf("Making a GET request to %v", path)

	response, err := http.Get(path)
	if err != nil {
		log.Fatalf("Error maiking the request %v", err)
	}

	return response, err
}

func ToStruct(response http.Response, target interface{}) error {
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}
