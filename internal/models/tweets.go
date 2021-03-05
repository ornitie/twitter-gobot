package models

type (
	TweetResponse struct {
		Tweet Tweet `json:"data" bson:"data"`
	}

	ResponseError struct {
		Errors []Error `json:"errors" bson:"errors"`
	}

	Tweet struct {
		ID     string `json:"id" bson:"id"`
		Text   string `json:"text" bson:"text"`
		Author string `json:"author_id" bson:"author_id"`
	}

	Error struct {
		Title   string `json:"title" bson:"title"`
		Type    string `json:"disconnect_type" bson:"disconnect_type"`
		Detail  string `json:"detail" bson:"detail"`
		TypeURL string `json:"type" bson:"type"`
	}
)
