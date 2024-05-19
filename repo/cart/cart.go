package cart

import (
	"context"
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

	obj := &entity.Cart{
		UserId: uId,
		BookId: bId,
	}
	tx := db.Model(obj).First(obj)

	if tx.RowsAffected > 0 {
		obj.Qty += qty
	} else {
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

func getBook(ctx context.Context, id uint) (*entity.Book, error) {
	db := database.GetDB(ctx)

	book := &entity.Book{Model: gorm.Model{ID: id}}

	err := db.Find(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}
