package resources

import (
	"net/http"
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

func (resource TweetsResource) StreamTweets() (*http.Response, error) {
	return resource.baseResource.Get(BASE_URL + STREAM_URI + AUTHOR_ID)
}
