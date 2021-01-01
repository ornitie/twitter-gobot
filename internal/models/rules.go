package models

type (
	RemoveRule struct {
		IDS []string `json:"ids,omitempty" bson:"ids"`
	}

	CreateRule struct {
		Add []Rule `json:"add" bson:"add" validate:"required"`
	}

	Rule struct {
		ID    string `json:"id" bson:"id"`
		Value string `json:"value" bson:"value" validate:"required"`
		Tag   string `json:"tag,omitempty" bson:"tag"`
	}

	RuleResponse struct {
		Rules    []Rule            `json:"data" bson:"data" validate:"required"`
		Metadata map[string]string `json:"meta,omitempty" bson:"meta"`
	}
)
