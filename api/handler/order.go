package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/service/order"
	"log/slog"
)

const orderHandlerName = "order"

type OrderHandler struct {
	svc *order.Service
}

var _ IHandler = (*OrderHandler)(nil)

func NewOrder(log *slog.Logger) IHandler {
	svc := order.New(log)

	return &OrderHandler{svc}
}

func (h OrderHandler) Name() string {
	return orderHandlerName
}

func (h OrderHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(h.Name())
	g.Post("/cart", h.AddToCart)
	g.Get("/cart/:user_id", h.GetCart)

	return h
}

func (h OrderHandler) AddToCart(c *fiber.Ctx) error {
	req := new(dto.OrderAddToCartRequest)

	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	cart, err := h.svc.AddToCart(c.Context(), req.UserId, req.BookId, req.Qty)
	if err != nil {
		return err
	}

	data := dto.CartFromEntity(cart)

	return middleware.OK(c, "item added to cart", data)
}

func (h OrderHandler) GetCart(c *fiber.Ctx) error {
	id, err := c.ParamsInt("user_id")
	if err != nil {
		return common.ErrInvalidUserId
	}

	cart, err := h.svc.GetChart(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.CartFromEntity(cart)

	return middleware.OK(c, "cart fetched", data)
}
