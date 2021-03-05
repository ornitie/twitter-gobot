package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ornitie/twitter-gobot/internal/models"
	"github.com/ornitie/twitter-gobot/internal/repositories"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type (
	TweetsService struct {
		tweetsResource *resources.TweetsResource
		baseRepository *repositories.TweetsRepository
	}
)

func NewTweetsService(baseResource *resources.BaseResource, baseRepository *repositories.BaseRepository) *TweetsService {
	return &TweetsService{tweetsResource: resources.NewTweetsResource(baseResource), baseRepository: repositories.NewTweetsRepository(baseRepository)}
}

func (service TweetsService) StreamTweets() {
	response, err := service.tweetsResource.StreamTweets()

	if err != nil {
		log.Fatalf("Error with the streaming request: %v", err)
	}

	for {
		responseTweet := &models.TweetResponse{}
		responseError := &models.ResponseError{}
		if response.Body != nil {
			reader := bufio.NewReader(response.Body)
			bytes, _ := reader.ReadBytes('\n')
			if len(bytes) > 2 {
				log.Printf("Message received %+v, stringified: %s", bytes, string(bytes))
				_ = json.Unmarshal(bytes, &responseError)
				if len(responseError.Errors) > 1 && responseError.Errors[0].Title != "" {
					service.StreamTweets()
					return
				}
				err := json.Unmarshal(bytes, &responseTweet)
				if err == nil {
					err = service.baseRepository.SaveTweet(responseTweet.Tweet)
					if err != nil {
						log.Printf("Error saving the tweet %+v, with error %v", responseTweet.Tweet, err)
					}
				} else {
					log.Printf("Error requesting the stream: %v", err)
				}
			}
		}
	}
}

func (service TweetsService) PrintStream() {
	response, err := service.tweetsResource.StreamTweets()

	if err != nil {
		log.Fatalf("Error with the streaming request: %v", err)
	}
	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Print("FUCK")
		}
		fmt.Printf("GOT ONE %s", string(line))
	}
}
