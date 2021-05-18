package infrastructure

import (
	"bufio"
	"context"
	"log"
	"os"

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

	var s string //TODO
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
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
	return nil, nil //TODO
}