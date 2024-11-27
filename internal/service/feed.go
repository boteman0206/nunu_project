package service

import (
    "context"
	"projectName/internal/model"
	"projectName/internal/repository"
)

type FeedService interface {
	GetFeed(ctx context.Context, id int64) (*model.Feed, error)
}
func NewFeedService(
    service *Service,
    feedRepository repository.FeedRepository,
) FeedService {
	return &feedService{
		Service:        service,
		feedRepository: feedRepository,
	}
}

type feedService struct {
	*Service
	feedRepository repository.FeedRepository
}

func (s *feedService) GetFeed(ctx context.Context, id int64) (*model.Feed, error) {
	return s.feedRepository.GetFeed(ctx, id)
}
