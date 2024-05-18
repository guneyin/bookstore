package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler"
	"log/slog"
)

type Api struct {
	handler *handler.Handler
}

func New(log *slog.Logger, router fiber.Router) *Api {
	hnd := handler.New(log, router)

	return &Api{handler: hnd}
}
