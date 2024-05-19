package handler

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type IHandler interface {
	Name() string
	SetRoutes(r fiber.Router) IHandler
}

type Handler struct {
	log      *slog.Logger
	router   fiber.Router
	handlers map[string]IHandler
}

func New(log *slog.Logger, router fiber.Router) *Handler {
	handler := &Handler{
		log:      log,
		router:   router,
		handlers: make(map[string]IHandler),
	}
	handler.registerHandlers()

	return handler
}

func (h Handler) registerHandlers() {
	h.registerHandler(NewGeneral)
	h.registerHandler(NewUser)
	h.registerHandler(NewBook)
	h.registerHandler(NewCart)
}

func (h Handler) registerHandler(f func(log *slog.Logger) IHandler) {
	hnd := f(h.log).SetRoutes(h.router)
	h.handlers[hnd.Name()] = hnd
}
