package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ornitie/twitter-gobot/internal/models"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type (
	TweetsService struct {
		tweetsResource *resources.TweetsResource
	}
)

func NewTweetsService(baseResource *resources.BaseResource) *TweetsService {
	return &TweetsService{tweetsResource: resources.NewTweetsResource(baseResource)}
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
			fmt.Printf("FUCK %v", err)
			return
		}

		fmt.Printf("YEP %+v", responseTweet.Tweet)
		//save()
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
