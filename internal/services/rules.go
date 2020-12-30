package services

import (
	"github.com/ornitie/twitter-gobot/internal/models"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

const (
	BASE_URL   = "https://api.twitter.com/2/"
	STREAM_URI = "tweets/search/stream/"
)

type (
	RulesService struct {
		baseResource *resources.BaseResource
	}
)

func NewRulesService(baseResource *resources.BaseResource) *RulesService {
	return &RulesService{baseResource: baseResource}
}

func (service RulesService) CreateNewRule(ruleValue string) error {
	createRule := models.CreateRule{Add: []models.Rule{models.Rule{Value: ruleValue}}}

	_, err := service.baseResource.Post(BASE_URL+STREAM_URI+"rules", createRule)

	return err
}

func (service RulesService) GetRules() (*models.RuleResponse, error) {
	response, _ := service.baseResource.Get("https://api.twitter.com/2/tweets/search/stream/rules")
	rulesResponse := &models.RuleResponse{}

	_ = resources.ToStruct(*response, rulesResponse)

	return rulesResponse, nil
}
