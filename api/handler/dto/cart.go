package dto

import "github.com/guneyin/bookstore/entity"

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
	BookId uint `json:"bookId"`
	Qty    uint `json:"qty"`
}

func CartFromEntity(ec []entity.CartResult) *CartResponse {
	if len(ec) == 0 {
		return nil
	}

	cart := &CartResponse{UserId: ec[0].ID}

	for _, item := range ec {
		cart.TotalPrice += item.TotalPrice()
		cart.Items = append(cart.Items, CartItem{
			BookId: item.BookId,
			Qty:    item.Qty,
		})
	}

	return cart
}
