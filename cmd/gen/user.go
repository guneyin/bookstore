package gen

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/repo/user"
	"gorm.io/gorm"
	"log/slog"
)

const urlUsers = "https://freetestapi.com/api/v1/users"

type UserDTO struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street string `json:"street"`
		City   string `json:"city"`
		Zip    string `json:"zip"`
	} `json:"address"`
	Phone string `json:"phone"`
}

func generateUserData(ctx context.Context, r *resty.Request, db *gorm.DB) error {
	log.InfoContext(ctx, "generating user data..")

	ul, err := fetchData[UserDTO](ctx, r, urlUsers)
	if err != nil {
		return err
	}

	truncateTable(db, &entity.User{})

	repo := user.NewRepo()
	for _, item := range ul {
		err = repo.Create(ctx, &entity.User{
			Name:     item.Name,
			Username: item.Username,
			Email:    item.Email,
			Address:  fmt.Sprintf("%s %s %s", item.Address.Street, item.Address.City, item.Address.Zip),
			Phone:    item.Phone,
		})
		if err != nil {
			log.WarnContext(ctx, fmt.Sprintf("%s could not created", item.Username), slog.String("err", err.Error()))
		}
	}

	return nil
}
