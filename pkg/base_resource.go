package resources

import (
	"log"
	"net/http"
)

func Get(path string) (*http.Response, error) {

	log.Printf("Making a GET request to %v", path)

	return http.Get(path)
}
