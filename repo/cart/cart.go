package cart

import (
	"context"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/database"
	"github.com/guneyin/bookstore/entity"
	"gorm.io/gorm"
)

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r Repo) AddToCart(ctx context.Context, uId, bId, qty uint) ([]entity.CartResult, error) {
	db := database.GetDB(ctx)

	_, err := getBook(ctx, bId)
	if err != nil {
		return nil, err
	}

	obj := &entity.Cart{}
	tx := db.First(obj, "user_id = ? and book_id = ?", uId, bId)

	if tx.RowsAffected > 0 {
		obj.Qty += qty
	} else {
		obj.BookId = bId
		obj.UserId = uId
		obj.Qty = qty
	}

	err = db.Save(obj).Error
	if err != nil {
		return nil, err
	}

	return r.GetCart(ctx, uId)
}

func (r Repo) GetCart(ctx context.Context, uId uint) ([]entity.CartResult, error) {
	db := database.GetDB(ctx)

	var cr []entity.CartResult

	obj := &entity.Cart{UserId: uId}
	err := db.Model(obj).Select("carts.*, books.price").Joins("inner join books on books.id = carts.book_id").Find(&cr).Error
	if err != nil {
		return nil, err
	}

	return cr, nil
}

func (r Repo) PlaceOrder(ctx context.Context, uId uint) (*entity.Order, error) {
	u, err := getUser(ctx, uId)
	if err != nil {
		return nil, err
	}

	sc, err := r.GetCart(ctx, uId)
	if err != nil {
		return nil, err
	}

	var orderPrice float64
	for _, item := range sc {
		orderPrice += item.TotalPrice()
	}

	db := database.GetDB(ctx).Begin()

	order := &entity.Order{
		UserId:  uId,
		Status:  common.OrderStatusCreated.ToInt(),
		Address: u.Address,
		Price:   orderPrice,
	}
	tx := db.Create(order)
	if tx.Error != nil {
		db.Rollback()

		return nil, err
	}

	var orderItems []entity.OrderItem
	for _, item := range sc {
		orderItems = append(orderItems, entity.OrderItem{
			OrderId:    order.ID,
			BookId:     item.BookId,
			Qty:        item.Qty,
			Price:      item.Price,
			TotalPrice: item.TotalPrice(),
		})
	}
	tx = db.Create(orderItems)
	if tx.Error != nil {
		db.Rollback()

		return nil, err
	}

	db.Commit()

	return r.GetOrder(ctx, order.ID)
}

func (r Repo) GetOrder(ctx context.Context, id uint) (*entity.Order, error) {
	db := database.GetDB(ctx)

	var (
		order      entity.Order
		orderItems []entity.OrderItem
	)

	err := db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	err = db.Find(&orderItems, "order_id", id).Error
	if err != nil {
		return nil, err
	}

	order.Items = orderItems

	return &order, nil
}

func (r Repo) GetOrdersByUserId(ctx context.Context, id uint) ([]entity.Order, error) {
	db := database.GetDB(ctx)

	var orders []entity.Order

	err := db.Find(&orders, "user_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func getBook(ctx context.Context, id uint) (*entity.Book, error) {
	db := database.GetDB(ctx)

	book := &entity.Book{Model: gorm.Model{ID: id}}

	err := db.Find(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func getUser(ctx context.Context, id uint) (*entity.User, error) {
	db := database.GetDB(ctx)

	obj := &entity.User{Model: gorm.Model{ID: id}}

	err := db.Find(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}
