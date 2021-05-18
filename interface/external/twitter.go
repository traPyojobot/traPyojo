package external

import "github.com/traPyojobot/traPyojo/domain"

type TwitterAPI interface {
	PostTweet(string) (*domain.Tweet, error)
}
