package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/dto"
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

func (gh *GeneralHandler) Name() string {
	return generalHandlerName
}

func (gh *GeneralHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(gh.Name())
	g.Get("status", gh.GeneralStatus)

	return gh
}

func (gh *GeneralHandler) GeneralStatus(c *fiber.Ctx) error {
	var status dto.StatusResponse
	status.FromEntity(gh.svc.Status())

	return middleware.OK(c, "service status fetched", status)
}
