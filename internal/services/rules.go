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
		rulesResource *resources.RulesResource
	}
)

func NewRulesService(baseResource *resources.BaseResource) *RulesService {
	return &RulesService{rulesResource: resources.NewRulesResource(baseResource)}
}

// func (service RulesService) CreateNewRule(ruleValue string) error {
// 	createRule := models.CreateRule{Add: []models.Rule{models.Rule{Value: ruleValue}}}

// 	_, err := service.baseResource.Post(BASE_URL+STREAM_URI+"rules", createRule)

// 	return err
// }

func (service RulesService) GetRules() (*models.RuleResponse, error) {
	response, error := service.rulesResource.GetRules()

	return response, error
}

func (service RulesService) DeleteRule(ruleId string) error {
	error := service.rulesResource.DeleteRule(ruleId)

	return error
}

func (service RulesService) CreateRule(createRule *models.CreateRule) error {
	error := service.rulesResource.CreateRule(createRule)

	return error
}
