package service

import (
	"context"

	"github.com/traPyojobot/traPyojo/domain"
	"github.com/traPyojobot/traPyojo/usecase/repository"
)

type TweetService struct {
	repo repository.TweetRepository
}

func NewTweetService(repo repository.TweetRepository) TweetService {
	return TweetService{
		repo,
	}
}

func (s *TweetService) GetMonologue(ctx context.Context) (*domain.Monologue, error) {
	return s.repo.GetMonologue()
}

func (s *TweetService) PostMonologue(ctx context.Context, content string) (*domain.Tweet, error) {
	return s.repo.PostMonologue(content)
}
