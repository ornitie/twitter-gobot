package controllers

import (
	"github.com/ornitie/twitter-gobot/internal/services"
	"github.com/ornitie/twitter-gobot/pkg/resources"
	"net/http"
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

func (controller *RulesController) GetRules(request *http.Request) interface{} {
	rules, _ := controller.service.GetRules()

	return rules
}
