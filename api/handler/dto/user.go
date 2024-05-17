package dto

import "github.com/guneyin/bookstore/repo/user"

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserListResponse []UserResponse

func UserFromEntity(u user.User) *UserResponse {
	return &UserResponse{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Name,
		Address:  u.Address,
		Phone:    u.Phone,
	}
}

func UserListFromEntity(ul *user.UserList) *UserListResponse {
	list := UserListResponse{}

	for _, item := range *ul {
		u := UserFromEntity(item)
		list = append(list, *u)
	}

	return &list
}
