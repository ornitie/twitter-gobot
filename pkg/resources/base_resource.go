package resources

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ornitie/twitter-gobot/pkg/models"
)

const (
	AUTH     = "Authorization"
	BEARER   = "Bearer "
	BASE_URL = "https://api.twitter.com/2/"
)

type BaseResource struct {
	bearer string
	client http.Client
}

func NewBaseResource(bearerToken string) *BaseResource {
	return &BaseResource{
		bearer: BEARER + bearerToken,
		client: http.Client{},
	}
}

func (br BaseResource) Get(path string) (*http.Response, error) {
	log.Printf("Making a GET request to %v", path)

	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Add(AUTH, br.bearer)

	response, err := br.client.Do(req)

	if err != nil {
		log.Fatalf("Error making the request %v", err)
	}

	return response, err
}

func (br BaseResource) Post(path string, payload interface{}) (*http.Response, error) {
	log.Printf("Making a POST request to %v", path)

	b, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", path, bytes.NewReader(b))

	req.Header.Add(AUTH, br.bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := br.client.Do(req)

	return response, err
}

func ToStruct(response http.Response, target interface{}) error {
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}

func NewResponseError(message string) models.ResourceError {
	return models.ResourceError{Message: message}
}
