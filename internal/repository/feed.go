package repository

import (
	"context"
	"projectName/internal/model"

	"gorm.io/gorm"
)

type FeedRepository interface {
	GetFeed(ctx context.Context, id int64) (*model.Feed, error)
	CreateFeed(ctx context.Context, feed *model.Feed) (int64, error)
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

// CreateFeed implements FeedRepository.
func (r *feedRepository) CreateFeed(ctx context.Context, feed *model.Feed) (int64, error) {

	result := r.db.WithContext(ctx).Create(&feed)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil

}

func (r *feedRepository) GetFeed(ctx context.Context, id int64) (*model.Feed, error) {
	var feed model.Feed

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&feed).Error
	if err == gorm.ErrRecordNotFound {
		return &feed, nil
	}
	if err != nil {
		return &feed, err
	}
	return &feed, nil
}
