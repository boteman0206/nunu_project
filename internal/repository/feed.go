package repository

import (
    "context"
	"projectName/internal/model"
)

type FeedRepository interface {
	GetFeed(ctx context.Context, id int64) (*model.Feed, error)
}

func NewFeedRepository(
	repository *Repository,
) FeedRepository {
	return &feedRepository{
		Repository: repository,
	}
}

type feedRepository struct {
	*Repository
}

func (r *feedRepository) GetFeed(ctx context.Context, id int64) (*model.Feed, error) {
	var feed model.Feed

	return &feed, nil
}
