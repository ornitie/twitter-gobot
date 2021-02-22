package models

type (
	RemoveRule struct {
		IDS []string `json:"ids,omitempty" bson:"ids"`
	}

	CreateRule struct {
		Add []Rule `json:"add" bson:"add" validate:"required"`
	}

	CreateResponse struct {
		Metadata Metadata            `json:"meta" bson:"meta"`
		Errors   []map[string]string `json:"errors" bson:"errors"`
		Data     []Rule              `json:"data" bson:"data"`
	}

	Rule struct {
		ID    string `json:"id,omitempty" bson:"id"`
		Value string `json:"value" bson:"value" validate:"required"`
		Tag   string `json:"tag,omitempty" bson:"tag"`
	}

	RuleResponse struct {
		Rules    []Rule            `json:"data" bson:"data" validate:"required"`
		Metadata map[string]string `json:"meta,omitempty" bson:"meta"`
	}

	Metadata struct {
		Sent    string           `json:"sent,omitempty" bson:"sent"`
		Summary map[string]int64 `json:"summary,omitempty" bson:"summary"`
	}
)
