package infrastructure

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traPyojobot/traPyojo/interface/handler"
)

func Init() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		return h(&handler.Context{Context: c})
	// 	}
	// })

	api, err := InjectAPIServer()
	if err != nil {
		log.Fatal(err)
	}

	echoApi := e.Group("/api")
	v1 := echoApi.Group("/v1")
	{
		v1.GET("/ping", handler.PingHandler)

		tweet := v1.Group("/tweet")
		{
			tweet.POST("/tweets/monologue", api.Tweet.PostMonologueHandler)
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
