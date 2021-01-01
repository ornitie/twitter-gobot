package resources

import (
	"github.com/ornitie/twitter-gobot/internal/models"
)

type RulesResource struct {
	baseResource *BaseResource
}

const (
	BASE_URL  = "https://api.twitter.com/2/"
	RULES_URI = "tweets/search/stream/rules"
)

func NewRulesResource(base *BaseResource) *RulesResource {
	return &RulesResource{
		baseResource: base,
	}
}

func (resource *RulesResource) GetRules() (*models.RuleResponse, error) {
	response, _ := resource.baseResource.Get(BASE_URL + RULES_URI)
	rulesResponse := &models.RuleResponse{}

	_ = ToStruct(*response, rulesResponse)

	return rulesResponse, nil
}

func (resource *RulesResource) DeleteRule(ID string) error {
	response, _ := resource.baseResource.Post(BASE_URL+RULES_URI, map[string]models.RemoveRule{
		"delete": models.RemoveRule{
			IDS: []string{ID},
		},
	})
	rulesResponse := &models.RuleResponse{}

	_ = ToStruct(*response, rulesResponse)

	return nil
}
