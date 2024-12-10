package service

import (
	"context"
	"projectName/internal/model"
	"projectName/internal/repository"
)

type FeedService interface {
	GetFeed(ctx context.Context, id int64) (*model.Feed, error)
	CreateFeed(ctx context.Context, feed *model.Feed) (int64, error)
}

func NewFeedService(
	service *Service,
	repository *repository.Repository,
	feedRepository repository.FeedRepository,
) FeedService {
	return &feedService{
		Service:        service,
		feedRepository: feedRepository,
		repository:     repository,
	}
}

type feedService struct {
	*Service
	feedRepository repository.FeedRepository
	repository     *repository.Repository
}

func (s *feedService) GetFeed(ctx context.Context, id int64) (*model.Feed, error) {

	return s.feedRepository.GetFeed(ctx, id)
}

func (s *feedService) CreateFeed(ctx context.Context, data *model.Feed) (int64, error) {

	return s.feedRepository.CreateFeed(ctx, data)
}
