package controllers

import (
	"github.com/ornitie/twitter-gobot/cmd/services",
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type (
	RulesController struct {
		services.RulesService
	}
)

func NewRulesController(baseResource resources.BaseResource) *RulesController {
	return &RulesController{
		RulesController: services.RulesService{
			SettingsService: services.NewRulesService(baseResource),
		},
	}
}
