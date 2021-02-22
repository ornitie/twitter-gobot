package controllers

import (
	"github.com/ornitie/twitter-gobot/internal/repositories"
	"github.com/ornitie/twitter-gobot/internal/services"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type (
	TweetsController struct {
		service services.TweetsService
	}
)

func NewTweetsController(baseResource *resources.BaseResource, baseRepository *repositories.BaseRepository) *TweetsController {
	return &TweetsController{
		service: *services.NewTweetsService(baseResource, baseRepository),
	}
}

func (controller TweetsController) StreamTweets() {
	controller.service.StreamTweets()
}
