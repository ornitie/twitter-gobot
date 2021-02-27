package repositories

import (
	"fmt"

	"github.com/ornitie/twitter-gobot/internal/models"
)

type TweetsRepository struct {
	baseRepository *BaseRepository
}

func NewTweetsRepository(baseRepository *BaseRepository) *TweetsRepository {
	return &TweetsRepository{baseRepository}
}

func (repository TweetsRepository) SaveTweet(tweet models.Tweet) error {
	query := fmt.Sprintf("insert into tweet(author_id, tweet_id, text) values('%s','%s','%s')", tweet.Author, tweet.ID, tweet.Text)

	return repository.baseRepository.Create(query)
}
