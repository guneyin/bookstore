package gen

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/repo/user"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type UserDTO struct {
	Id       int    `json:"id"`
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

const urlUserList = "https://freetestapi.com/api/v1/users"

func generateUserData(ctx context.Context, r *resty.Request, db *gorm.DB) error {
	common.Log.InfoContext(ctx, "generating user data..")
	var ul []UserDTO

	res, err := r.
		SetResult(&ul).
		SetQueryParam("limit", "10").
		Get(urlUserList)
	if err != nil {
		common.Log.ErrorContext(ctx, "error on api call", slog.String("msg", err.Error()))

		return err
	}

	if res.StatusCode() >= http.StatusBadRequest {
		common.Log.ErrorContext(ctx, "api returned error",
			slog.Int("status_code", res.StatusCode()),
			slog.String("status", res.Status()),
		)

		return errors.New(res.Status())
	}

	truncateTable(db, &user.User{})

	for _, entity := range ul {
		err = user.Create(ctx, &user.User{
			Id:       entity.Id,
			Name:     entity.Name,
			Username: entity.Username,
			Email:    entity.Email,
			Address:  fmt.Sprintf("%s %s %s", entity.Address.Street, entity.Address.City, entity.Address.Zip),
			Phone:    entity.Phone,
		})
		if err != nil {
			common.Log.WarnContext(ctx, fmt.Sprintf("%s could not created", entity.Username), slog.String("err", err.Error()))
		}
	}

	return nil
}
