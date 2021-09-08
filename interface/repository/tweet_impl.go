package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/traPyojobot/traPyojo/domain"
	"github.com/traPyojobot/traPyojo/interface/external"
	"github.com/traPyojobot/traPyojo/usecase/repository"
)

var (
	GatewayEndpoint = os.Getenv("GatewayEndpoint")
)

type TweetRepository struct {
	api external.TwitterAPI
}

func NewTweetRepository() repository.TweetRepository {
	return &TweetRepository{}
}

func (repo *TweetRepository) GetMonologue() (*domain.Monologue, error) {
	resp, err := http.Get(fmt.Sprintf("%s/monologue", GatewayEndpoint))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	monologue := &domain.Monologue{}
	if err := json.Unmarshal(body, monologue); err != nil {
		return nil, err
	}

	return monologue, nil
}

func (repo *TweetRepository) PostMonologue(content string) (*domain.Tweet, error) {
	res, err := repo.api.PostTweet(content)
	if err != nil {
		return nil, err
	}

	return res, nil
}
