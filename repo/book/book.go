package book

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

func (r Repo) Create(ctx context.Context, u *entity.Book) error {
	db := database.DB.WithContext(ctx)

	b := &entity.Book{}
	db.Model(u).First(r)

	if b.Title == u.Title {
		return common.ErrAlreadyExist
	}

	return db.Save(u).Error
}

func (r Repo) GetList(ctx context.Context) (*entity.BookList, error) {
	db := database.DB.WithContext(ctx)

	var bl entity.BookList
	err := db.Model(&entity.Book{}).Find(&bl).Error
	if err != nil {
		return nil, err
	}

	return &bl, nil
}

func (r Repo) GetById(ctx context.Context, id uint) (*entity.Book, error) {
	db := database.DB.WithContext(ctx)

	u := &entity.Book{Model: gorm.Model{ID: id}}
	err := db.Model(u).First(u).Error
	if err != nil {
		return nil, common.ErrNotFound
	}

	return u, nil
}

func (r Repo) Search(ctx context.Context, sp *entity.BookSearchParams) (*entity.BookList, error) {
	db := database.DB.WithContext(ctx)

	var bl entity.BookList
	err := db.Where("title like ? and author like ? and genre like ?", sp.Title.ToWildcard(), sp.Author.ToWildcard(), sp.Genre.ToWildcard()).
		Find(&bl).Error
	if err != nil {
		return nil, err
	}

	return &bl, nil
}
