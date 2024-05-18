package dto

import (
	"github.com/guneyin/bookstore/entity"
)

type UserResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserListResponse []UserResponse

func UserFromEntity(u entity.User) *UserResponse {
	return &UserResponse{
		Id:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Name,
		Address:  u.Address,
		Phone:    u.Phone,
	}
}

func UserListFromEntity(ul *entity.UserList) *UserListResponse {
	list := UserListResponse{}

	for _, item := range *ul {
		u := UserFromEntity(item)
		list = append(list, *u)
	}

	return &list
}
