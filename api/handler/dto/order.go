package dto

import "github.com/guneyin/bookstore/entity"

type OrderAddToCartRequest struct {
	UserId uint `json:"userId"`
	BookId uint `json:"bookId"`
	Qty    uint `json:"qty"`
}

type OrderCartResponse struct {
	UserId     uint            `json:"userId"`
	TotalPrice float64         `json:"totalPrice"`
	Items      []OrderCartItem `json:"items"`
}

type OrderCartItem struct {
	BookId uint `json:"bookId"`
	Qty    uint `json:"qty"`
}

func CartFromEntity(ec []entity.CartResult) *OrderCartResponse {
	if len(ec) == 0 {
		return nil
	}

	cart := &OrderCartResponse{UserId: ec[0].ID}

	for _, item := range ec {
		cart.TotalPrice += item.TotalPrice()
		cart.Items = append(cart.Items, OrderCartItem{
			BookId: item.BookId,
			Qty:    item.Qty,
		})
	}

	return cart
}
