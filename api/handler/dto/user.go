package dto

import "github.com/guneyin/bookstore/service/user"

type UserResponse struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Address  UserAddress `json:"address"`
	Phone    string      `json:"phone"`
}

type UserAddress struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type UserListResponse []UserResponse

func UserFromEntity(u user.User) *UserResponse {
	return &UserResponse{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Name,
		Address: UserAddress{
			Street:  u.Address.Street,
			Suite:   u.Address.Suite,
			City:    u.Address.City,
			Zipcode: u.Address.Zipcode,
		},
		Phone: u.Phone,
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
