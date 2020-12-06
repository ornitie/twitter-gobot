package services

import (
	"github.com/ornitie/twitter-gobot/cmd/models"
	"github.com/ornitie/twitter-gobot/pkg"
)

const (
	BASE_URL   = "https://api.twitter.com/2/"
	STREAM_URI = "tweets/search/stream/"
)

type (
	RulesService struct {
		baseResource resources.BaseResource
	}
)

func NewRulesService(baseResource resources.BaseResource) *RulesService {
	return &RulesService{baseResource: baseResource}
}

func (service RulesService) CreateNewRule(ruleValue string) error {
	createRule := models.CreateRule{Add: []models.Rule{models.Rule{Value: ruleValue}}}

	_, err := service.baseResource.Post(BASE_URL+STREAM_URI+"rules", createRule)

	return err
}
