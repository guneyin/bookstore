package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/config"
)

type IHandler interface {
	SetRoutes(r fiber.Router) IHandler
}

type Handler struct {
	cfg      *config.Config
	handlers []IHandler
}

func New(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) InitHandlers(r fiber.Router) {
	h.handlers = []IHandler{
		NewGeneral(h.cfg).SetRoutes(r),
	}
}
