package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/service/user"
	"log/slog"
)

const userHandlerName = "user"

type UserHandler struct {
	svc *user.Service
}

var _ IHandler = (*UserHandler)(nil)

func NewUser(log *slog.Logger) IHandler {
	svc := user.New(log)

	return &UserHandler{svc}
}

func (h UserHandler) Name() string {
	return userHandlerName
}

func (h UserHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(h.Name())
	g.Get("list", h.GetUserList)
	g.Get("/:id", h.GetUserById)

	return h
}

// GetUserList
// @Summary List users.
// @Description List all the users.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} middleware.ResponseHTTP{data=dto.UserListResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /user/list [get]
func (h UserHandler) GetUserList(c *fiber.Ctx) error {
	obj, err := h.svc.GetList(c.Context())
	if err != nil {
		return err
	}

	data := dto.UserListFromEntity(obj)

	return middleware.OK(c, fmt.Sprintf("%d users fetched", len(*data)), data)
}

// GetUserById
// @Summary Get user.
// @Description Get user by id.
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.UserResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /user/{id} [get]
func (h UserHandler) GetUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id", 0)

	obj, err := h.svc.GetById(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.UserFromEntity(obj)

	return middleware.OK(c, "user fetched", data)
}
