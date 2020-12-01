package resources

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	AUTH   = "Authorization"
	BEARER = "Bearer "
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
		log.Fatalf("Error maiking the request %v", err)
	}

	return response, err
}

func ToStruct(response http.Response, target interface{}) error {
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}
