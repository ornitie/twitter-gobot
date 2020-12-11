package models

type (
	CreateRule struct {
		Add []Rule `json:"add" bson:"add" validate:"required"`
	}

	Rule struct {
		ID    string `json:"id" bson:"id"`
		Value string `json:"value" bson:"value" validate:"required"`
		Tag   string `json:"tag,omitempty" bson:"tag"`
	}
)
