package infrastructure

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/traPyojobot/traPyojo/domain"
	"github.com/traPyojobot/traPyojo/interface/external"
	"golang.org/x/oauth2"
)

type TwitterAPI struct {
	Client *twitter.Client
}

func NewTwitterAPI() (external.TwitterAPI, error) {
	config := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	ctx := context.Background()

	var s string
	if sc := bufio.NewScanner(os.Stdin); sc.Scan() {
		s = sc.Text()
	}

	token, err := config.Exchange(ctx, s)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := config.Client(ctx, token)

	client := twitter.NewClient(httpClient)

	return &TwitterAPI{Client: client}, nil
}

func (twitter *TwitterAPI) PostTweet(content string) (*domain.Tweet, error) {
	tweet, resp, err := twitter.Client.Statuses.Update(content, nil)
	if err != nil {
		return &domain.Tweet{Response: resp}, err
	}

	ca, err := ConvertToTime(tweet.CreatedAt)
	if err != nil {
		return &domain.Tweet{Response: resp}, err
	}

	res := &domain.Tweet{
		ID:        tweet.ID,
		Response:  resp,
		Content:   tweet.FullText,
		CreatedAt: ca,
	}

	return res, err
}

func ConvertToTime(str string) (time.Time, error) {
	t, err := time.Parse(str, str)
	if err != nil {
		return time.Time{}, errors.New("invalid timestamp")
	}

	return t, nil
}
