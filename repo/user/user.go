package user

import (
	"context"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/database"
)

func Create(ctx context.Context, u *User) error {
	db := database.DB.WithContext(ctx)

	r := &User{}
	db.Model(u).First(r)

	if r.Id == u.Id {
		return common.ErrAlreadyExist
	}

	return db.Save(u).Error
}

func GetList(ctx context.Context) (UserList, error) {
	db := database.DB.WithContext(ctx)

	var ul UserList
	err := db.Model(&User{}).Find(&ul).Error
	if err != nil {
		return nil, err
	}

	return ul, nil
}

func GetById(ctx context.Context, id int) (*User, error) {
	db := database.DB.WithContext(ctx)

	u := &User{Id: id}
	err := db.Model(u).First(u).Error
	if err != nil {
		return nil, common.ErrNotFound
	}

	return u, nil
}
