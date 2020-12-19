package controllers

import (
	"github.com/ornitie/twitter-gobot/cmd/services"
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
