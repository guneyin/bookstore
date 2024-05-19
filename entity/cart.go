package entity

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type Cart struct {
	gorm.Model
	UserId uint `gorm:"uniqueIndex:idx_cart"`
	BookId uint `gorm:"uniqueIndex:idx_cart"`
	Qty    uint
}

type CartResult struct {
	Cart
	Price float64
}

type Order struct {
	gorm.Model
	UserId  uint `gorm:"index"`
	Status  uint
	Address string
	Price   float64
	Items   []OrderItem `gorm:"-"`
}

type OrderItem struct {
	gorm.Model
	OrderId    uint `gorm:"index"`
	BookId     uint
	Qty        uint
	Price      float64
	TotalPrice float64
}

func (cr CartResult) TotalPrice() float64 {
	s := fmt.Sprintf("%.2f", cr.Price*float64(cr.Qty))

	p, _ := strconv.ParseFloat(s, 64)
	return p
}

func (o Order) TotalPrice() float64 {
	var tp float64

	for _, item := range o.Items {
		tp += item.TotalPrice
	}

	s := fmt.Sprintf("%.2f", tp)

	p, _ := strconv.ParseFloat(s, 64)
	return p
}
