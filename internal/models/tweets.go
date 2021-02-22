package models

type (
	TweetResponse struct {
		Tweet Tweet `json:"data" bson:"data"`
	}

	Tweet struct {
		ID     string `json:"id" bson:"id"`
		Text   string `json:"text" bson:"text"`
		Author string `json:"author_id" bson:"author_id"`
	}
)
