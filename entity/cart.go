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

func (cr CartResult) TotalPrice() float64 {
	s := fmt.Sprintf("%.2f", cr.Price*float64(cr.Qty))

	p, _ := strconv.ParseFloat(s, 64)
	return p
}
