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
		err := json.NewDecoder(response.Body).Decode(&responseTweet)
		if err != nil {
			log.Fatalf("Error requesting the stream: %v", err)
		}

		fmt.Printf("YEP %+v", responseTweet.Tweet)
		err = service.baseRepository.SaveTweet(responseTweet.Tweet)
		if err != nil {
			log.Printf("Error saving the tweet %+v, with error %v", responseTweet.Tweet, err)
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
