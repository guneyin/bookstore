package user

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/config"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

const (
	urlUserList = "https://jsonplaceholder.typicode.com/users"
)

type Service struct {
	cfg   *config.Config
	httpC *resty.Client
}

func New(cfg *config.Config) *Service {
	return &Service{
		cfg:   cfg,
		httpC: resty.New(),
	}
}

func (s Service) GetUserList(ctx context.Context) (*UserList, error) {
	common.Log.InfoContext(ctx, "entered GetUserList")

	var list UserList

	res, err := s.httpC.R().
		SetResult(&list).
		Get(urlUserList)
	if err != nil {
		common.Log.ErrorContext(ctx, "error on api call", slog.String("msg", err.Error()))

		return nil, err
	}

	if res.StatusCode() >= http.StatusBadRequest {
		common.Log.ErrorContext(ctx, "api returned error",
			slog.Int("status_code", res.StatusCode()),
			slog.String("status", res.Status()),
		)

		return nil, errors.New(res.Status())
	}

	common.Log.InfoContext(ctx, "user list fetched", slog.Int("count", len(list)))

	return &list, nil
}

func (s Service) GetUserById(ctx context.Context, id int) (*User, error) {
	common.Log.InfoContext(ctx, "entered GetUserById")

	var user User

	u, _ := url.JoinPath(urlUserList, strconv.Itoa(id))

	res, err := s.httpC.R().
		SetResult(&user).
		Get(u)
	if err != nil {
		common.Log.ErrorContext(ctx, "error on api call", slog.String("msg", err.Error()))

		return nil, err
	}

	if res.StatusCode() >= http.StatusBadRequest {
		common.Log.ErrorContext(ctx, "api returned error",
			slog.Int("status_code", res.StatusCode()),
			slog.String("status", res.Status()),
		)

		return nil, errors.New(res.Status())
	}

	common.Log.InfoContext(ctx, "user fetched")

	return &user, nil
}
