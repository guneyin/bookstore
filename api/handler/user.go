package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/config"
	"github.com/guneyin/bookstore/service/user"
)

const userHandlerName = "user"

type UserHandler struct {
	svc *user.Service
}

var _ IHandler = (*UserHandler)(nil)

func NewUser(cfg *config.Config) IHandler {
	svc := user.New(cfg)

	return &UserHandler{svc}
}

func (h *UserHandler) Name() string {
	return userHandlerName
}

func (h *UserHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(h.Name())
	g.Get("list", h.GetUserList)
	g.Get("/:id", h.GetUserById)

	return h
}

func (h *UserHandler) GetUserList(c *fiber.Ctx) error {
	list, err := h.svc.GetUserList(c.Context())
	if err != nil {
		return err
	}

	data := dto.UserListFromEntity(list)

	return middleware.OK(c, "user list fetched", data)
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id", 0)

	data, err := h.svc.GetUserById(c.Context(), id)
	if err != nil {
		return err
	}

	return middleware.OK(c, "user fetched", data)
}
