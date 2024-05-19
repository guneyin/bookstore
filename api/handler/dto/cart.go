package dto

import (
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/entity"
	"time"
)

type AddToCartRequest struct {
	UserId uint `json:"userId"`
	BookId uint `json:"bookId"`
	Qty    uint `json:"qty"`
}

type CartResponse struct {
	UserId     uint       `json:"userId"`
	TotalPrice float64    `json:"totalPrice"`
	Items      []CartItem `json:"items"`
}

type CartItem struct {
	BookId     uint    `json:"bookId"`
	Qty        uint    `json:"qty"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"totalPrice"`
}

type OrderResponse struct {
	Id      uint      `json:"id"`
	Date    time.Time `json:"date"`
	UserId  uint      `json:"userId"`
	Price   float64   `json:"price"`
	Address string    `json:"address"`
	Status  string    `json:"status"`
}

type PlaceOrderResponse struct {
	OrderResponse
	Items []OrderItem `json:"items"`
}

type UserOrdersResponse struct {
	UserId uint            `json:"userId"`
	Orders []OrderResponse `json:"orders"`
}

type OrderItem struct {
	BookId     uint    `json:"bookId"`
	Qty        uint    `json:"qty"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"totalPrice"`
}

func CartFromEntity(ec []entity.CartResult) *CartResponse {
	if len(ec) == 0 {
		return nil
	}

	cart := &CartResponse{UserId: ec[0].Cart.UserId}

	for _, item := range ec {
		cart.TotalPrice += item.TotalPrice()
		cart.Items = append(cart.Items, CartItem{
			BookId:     item.BookId,
			Qty:        item.Qty,
			Price:      item.Price,
			TotalPrice: item.TotalPrice(),
		})
	}

	return cart
}

func OrderFromEntity(o *entity.Order) *PlaceOrderResponse {
	order := &PlaceOrderResponse{
		OrderResponse: OrderResponse{
			Id:      o.ID,
			Date:    o.CreatedAt,
			UserId:  o.UserId,
			Price:   o.TotalPrice(),
			Address: o.Address,
			Status:  common.OrderStatus(o.Status).ToString(),
		},
	}

	for _, item := range o.Items {
		order.Items = append(order.Items, OrderItem{
			BookId:     item.BookId,
			Qty:        item.Qty,
			Price:      item.Price,
			TotalPrice: item.TotalPrice,
		})
	}

	return order
}

func UserOrdersFromEntity(userId uint, orders []entity.Order) *UserOrdersResponse {
	res := &UserOrdersResponse{
		UserId: userId,
	}

	for _, item := range orders {
		res.Orders = append(res.Orders, OrderResponse{
			Id:      item.ID,
			Date:    item.CreatedAt,
			UserId:  item.UserId,
			Price:   item.Price,
			Address: item.Address,
			Status:  common.OrderStatus(item.Status).ToString(),
		})
	}

	return res
}
