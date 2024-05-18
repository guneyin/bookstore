package user

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

func (r Repo) Create(ctx context.Context, u *entity.User) error {
	db := database.GetDB(ctx)

	obj := &entity.User{}

	db.Model(&entity.User{}).First(obj)

	if obj.Email == u.Email {
		return common.ErrAlreadyExist
	}

	return db.Save(u).Error
}

func (r Repo) GetList(ctx context.Context) (*entity.UserList, error) {
	db := database.GetDB(ctx)

	var ul *entity.UserList
	err := db.Model(&entity.User{}).Find(ul).Error
	if err != nil {
		return nil, err
	}

	return ul, nil
}

func (r Repo) GetById(ctx context.Context, id uint) (*entity.User, error) {
	db := database.GetDB(ctx)

	u := &entity.User{Model: gorm.Model{ID: id}}
	err := db.Model(u).First(u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}
