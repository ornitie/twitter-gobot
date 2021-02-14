package controllers

import (
	"github.com/ornitie/twitter-gobot/internal/services"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type (
	TweetsController struct {
		service services.TweetsService
	}
)

func NewTweetsController(baseResource *resources.BaseResource) *TweetsController {
	return &TweetsController{
		service: *services.NewTweetsService(baseResource),
	}
}

func (controller TweetsController) StreamTweets() {
	controller.service.StreamTweets()
}
