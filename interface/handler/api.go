package handler

type API struct {
	Tweet *TweetHandler
}

func NewAPI(tweet *TweetHandler) API {
	return API{
		Tweet: tweet,
	}
}
