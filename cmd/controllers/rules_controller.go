package controllers

import (
	"github.com/gorilla/mux"
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

func (controller *RulesController) DeleteRule(request *http.Request) interface{} {
	ruleId := mux.Vars(request)["ID"]
	error := controller.service.DeleteRule(ruleId)

	response := map[string]string{
		"status": "Rule removed succesfully",
	}

	if error != nil {
		response["status"] = "Error removing Rule"
		response["error"] = error.Error()
	}

	return response
}
