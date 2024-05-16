package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler"
	"github.com/guneyin/bookstore/config"
)

type Api struct {
	handler *handler.Handler
}

func New(cfg *config.Config, router fiber.Router) *Api {
	hnd := handler.New(cfg, router)

	return &Api{handler: hnd}
}
