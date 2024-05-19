package user

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

func (r Repo) Create(ctx context.Context, u *entity.User) error {
	db := database.GetDB(ctx)

	obj := &entity.User{}
	tx := db.First(obj, "email = ?", u.Email)

	if tx.RowsAffected == 1 {
		return common.ErrAlreadyExist
	}

	return db.Save(u).Error
}

func (r Repo) GetList(ctx context.Context) (*entity.UserList, error) {
	db := database.GetDB(ctx)

	obj := &entity.UserList{}
	err := db.Model(obj).Find(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repo) GetById(ctx context.Context, id uint) (*entity.User, error) {
	db := database.GetDB(ctx)

	obj := &entity.User{}
	err := db.First(obj, id).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}
