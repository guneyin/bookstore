package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/service/cart"
	"log/slog"
)

const cartHandlerName = "cart"

type CartHandler struct {
	svc *cart.Service
}

var _ IHandler = (*CartHandler)(nil)

func NewCart(log *slog.Logger) IHandler {
	svc := cart.New(log)

	return &CartHandler{svc}
}

func (h CartHandler) Name() string {
	return cartHandlerName
}

func (h CartHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(h.Name())
	g.Post("/", h.Add)
	g.Get("/:user_id", h.GetCartByUserId)
	g.Get("/place-order/:user_id", h.PlaceOrder)
	g.Get("/order/:order_id", h.GetOrderById)
	g.Get("/order/user/:user_id", h.GetOrdersByUserId)

	return h
}

// Add
// @Summary Add to cart.
// @Description Add book to user cart.
// @Tags cart
// @Accept json
// @Produce json
// @Param search body dto.AddToCartRequest true "Add to cart"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.CartResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /cart [post]
func (h CartHandler) Add(c *fiber.Ctx) error {
	req := new(dto.AddToCartRequest)

	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	sc, err := h.svc.AddToCart(c.Context(), req.UserId, req.BookId, req.Qty)
	if err != nil {
		return err
	}

	data := dto.CartFromEntity(sc)

	return middleware.OK(c, "item added to cart", data)
}

// GetCartByUserId
// @Summary Get user cart.
// @Description Get user cart.
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.CartResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /cart/{id} [get]
func (h CartHandler) GetCartByUserId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("user_id")
	if err != nil {
		return common.ErrInvalidUserId
	}

	sc, err := h.svc.GetCart(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.CartFromEntity(sc)

	return middleware.OK(c, "cart fetched", data)
}

// PlaceOrder
// @Summary Place order.
// @Description Order all items in the user cart.
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.PlaceOrderResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /cart/place-order/{id} [get]
func (h CartHandler) PlaceOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("user_id")
	if err != nil {
		return common.ErrInvalidUserId
	}

	order, err := h.svc.PlaceOrder(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.OrderFromEntity(order)

	return middleware.OK(c, "order created", data)
}

// GetOrderById
// @Summary Get order.
// @Description Get order by id.
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.PlaceOrderResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /cart/order/{id} [get]
func (h CartHandler) GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("order_id")
	if err != nil {
		return common.ErrInvalidUserId
	}

	order, err := h.svc.GetOrderById(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.OrderFromEntity(order)

	return middleware.OK(c, "order fetched", data)
}

// GetOrdersByUserId
// @Summary Get user orders.
// @Description Get orders by user id.
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.UserOrdersResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /cart/order/user/{id} [get]
func (h CartHandler) GetOrdersByUserId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("user_id")
	if err != nil {
		return common.ErrInvalidUserId
	}

	orders, err := h.svc.GetOrdersByUserId(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.UserOrdersFromEntity(uint(id), orders)

	return middleware.OK(c, "user orders fetched", data)
}
