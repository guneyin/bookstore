package user

import (
	"context"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/repo/user"
	"log/slog"
)

type Service struct {
	repo *user.Repo
	log  *slog.Logger
}

func New(log *slog.Logger) *Service {
	return &Service{
		repo: user.NewRepo(),
		log:  log,
	}
}

func (s Service) GetList(ctx context.Context) (*entity.UserList, error) {
	s.log.InfoContext(ctx, "entered GetList")

	ul, err := s.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}

	s.log.InfoContext(ctx, "user list fetched", slog.Int("count", len(*ul)))

	return ul, nil
}

func (s Service) GetById(ctx context.Context, id uint) (*entity.User, error) {
	s.log.InfoContext(ctx, "entered GetById")

	u, err := s.repo.GetById(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on GetById", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "user fetched")

	return u, nil
}

func (s Service) Search(ctx context.Context, sp *entity.BookSearchParams) (*entity.BookList, error) {
	return nil, nil
}
