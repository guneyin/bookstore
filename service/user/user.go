package user

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/config"
	"github.com/guneyin/bookstore/repo/user"
	"log/slog"
)

//

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

func (s Service) GetUserList(ctx context.Context) (*user.UserList, error) {
	common.Log.InfoContext(ctx, "entered GetUserList")

	ul, err := user.GetList(ctx)
	if err != nil {
		return nil, err
	}

	common.Log.InfoContext(ctx, "user list fetched", slog.Int("count", len(ul)))

	return &ul, nil
}

func (s Service) GetUserById(ctx context.Context, id int) (*user.User, error) {
	common.Log.InfoContext(ctx, "entered GetUserById")

	u, err := user.GetById(ctx, id)
	if err != nil {
		common.Log.ErrorContext(ctx, "error on GetUserById", slog.String("msg", err.Error()))

		return nil, err
	}

	common.Log.InfoContext(ctx, "user fetched")

	return u, nil
}
