//+build wireinject

package infrastructure

import (
	"github.com/google/wire"
	"github.com/traPyojobot/traPyojo/interface/handler"
	impl "github.com/traPyojobot/traPyojo/interface/repository"
	"github.com/traPyojobot/traPyojo/usecase/service"
)

var tweetSet = wire.NewSet(
	NewTwitterAPI,
	impl.NewTweetRepository,
	service.NewTweetService,
)

var apiSet = wire.NewSet(
	handler.NewTweetHandler,
	handler.NewAPI,
)

func InjectAPIServer() (handler.API, error) {
	wire.Build(
		tweetSet,
		apiSet,
	)
	return handler.API{}, nil
}
