package resources

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/ornitie/twitter-gobot/internal/models"
)

const (
	STREAM_URI = "tweets/search/stream?"
	AUTHOR_ID  = "tweet.fields=author_id&"
)

type TweetsResource struct {
	baseResource *BaseResource
}

func NewTweetsResource(base *BaseResource) *TweetsResource {
	return &TweetsResource{
		baseResource: base,
	}
}

func (resource TweetsResource) StreamTweets() {
	response, _ := resource.baseResource.Get(BASE_URL + STREAM_URI + AUTHOR_ID)
	for {
		responseTweet := &models.TweetResponse{}
		err := json.NewDecoder(response.Body).Decode(&responseTweet)
		if err != nil {
			fmt.Printf("FUCK %v", err)
			return
		}

		fmt.Printf("YEP %+v", responseTweet.Tweet)
	}
}

func (resource TweetsResource) PrintString() {
	response, _ := resource.baseResource.Get(BASE_URL + STREAM_URI + AUTHOR_ID)
	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Print("FUCK")
		}
		fmt.Printf("GOT ONE %s", string(line))
	}
}
