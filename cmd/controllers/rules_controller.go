package controllers

import (
	"github.com/ornitie/twitter-gobot/internal/models"
	"github.com/ornitie/twitter-gobot/internal/services"
	"github.com/ornitie/twitter-gobot/pkg/resources"
)

type (
	RulesController struct {
		service services.RulesService
	}
)

func NewRulesController(baseResource *resources.BaseResource) *RulesController {
	return &RulesController{
		service: *services.NewRulesService(baseResource),
	}
}

func (controller *RulesController) GetRules() *models.RuleResponse {
	rules, _ := controller.service.GetRules()

	return rules
}
