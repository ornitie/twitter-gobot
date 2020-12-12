package services

import "github.com/ornitie/twitter-gobot/pkg"

type (
	RulesService struct {
		baseResource resources.BaseResource
	}
)

func NewRulesService(baseResource resources.BaseResource) *RulesService {
	return &RulesService{baseResource: baseResource}
}
