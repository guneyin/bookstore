package book

import (
	"context"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/database"
	"github.com/guneyin/bookstore/entity"
)

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r Repo) Create(ctx context.Context, u *entity.Book) error {
	db := database.GetDB(ctx)

	obj := &entity.Book{}
	tx := db.First(obj, "title = ?", u.Title)
	if tx.RowsAffected == 1 {
		return common.ErrAlreadyExist
	}

	return db.Create(u).Error
}

func (r Repo) GetList(ctx context.Context) (*entity.BookList, error) {
	db := database.GetDB(ctx)

	obj := &entity.BookList{}
	err := db.Find(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repo) GetById(ctx context.Context, id uint) (*entity.Book, error) {
	db := database.GetDB(ctx)

	obj := &entity.Book{}
	err := db.First(obj, id).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repo) Search(ctx context.Context, sp *entity.BookSearchParams) (*entity.BookList, error) {
	db := database.GetDB(ctx)

	obj := &entity.BookList{}
	err := db.Where("title like ? and author like ? and genre like ?", sp.Title.ToWildcard(), sp.Author.ToWildcard(), sp.Genre.ToWildcard()).
		Find(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}
