package services

import (
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
