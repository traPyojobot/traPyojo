package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traPyojobot/traPyojo/usecase/service"
)

type TweetHandler struct {
	srv service.TweetService
}

func NewTweetHandler(srv service.TweetService) *TweetHandler {
	return &TweetHandler{srv}
}

type MonologueResponse struct {
	Content string `json:"content"`
}

func (h *TweetHandler) PostMonologueHandler(c echo.Context) error {
	ctx := c.Request().Context()

	monologue, err := h.srv.GetMonologue(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := h.srv.PostMonologue(ctx, monologue.Content)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
