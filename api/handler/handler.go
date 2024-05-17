package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/config"
)

type IHandler interface {
	Name() string
	SetRoutes(r fiber.Router) IHandler
}

type Handler struct {
	cfg      *config.Config
	router   fiber.Router
	handlers map[string]IHandler
}

func New(cfg *config.Config, router fiber.Router) *Handler {
	handler := &Handler{
		cfg:      cfg,
		router:   router,
		handlers: make(map[string]IHandler),
	}
	handler.registerHandlers()

	return handler
}

func (h *Handler) registerHandlers() {
	h.registerHandler(NewGeneral)
	h.registerHandler(NewUser)
}

func (h *Handler) registerHandler(f func(cfg *config.Config) IHandler) {
	hnd := f(h.cfg).SetRoutes(h.router)
	h.handlers[hnd.Name()] = hnd
}
