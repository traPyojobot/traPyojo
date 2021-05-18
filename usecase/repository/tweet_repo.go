package repository

import "github.com/traPyojobot/traPyojo/domain"

type TweetRepository interface {
	GetMonologue() (*domain.Monologue, error)
	PostMonologue(string) (*domain.Tweet, error)
}
