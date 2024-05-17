package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/config"
	"github.com/guneyin/bookstore/service/general"
)

const generalHandlerName = "general"

type GeneralHandler struct {
	svc *general.Service
}

var _ IHandler = (*GeneralHandler)(nil)

func NewGeneral(cfg *config.Config) IHandler {
	svc := general.New(cfg)

	return &GeneralHandler{svc}
}

func (h *GeneralHandler) Name() string {
	return generalHandlerName
}

func (h *GeneralHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(h.Name())
	g.Get("status", h.GeneralStatus)

	return h
}

func (h *GeneralHandler) GeneralStatus(c *fiber.Ctx) error {
	status := dto.StatusFromEntity(h.svc.Status())

	return middleware.OK(c, "service status fetched", status)
}
