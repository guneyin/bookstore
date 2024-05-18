package book

import (
	"context"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/repo/book"
	"log/slog"
)

type Service struct {
	repo *book.Repo
	log  *slog.Logger
}

func New(log *slog.Logger) *Service {
	return &Service{
		repo: book.NewRepo(),
		log:  log,
	}
}

func (s Service) GetList(ctx context.Context) (*entity.BookList, error) {
	s.log.InfoContext(ctx, "entered GetList")

	bl, err := s.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}

	s.log.InfoContext(ctx, "book list fetched", slog.Int("count", len(*bl)))

	return bl, nil
}

func (s Service) GetById(ctx context.Context, id uint) (*entity.Book, error) {
	s.log.InfoContext(ctx, "entered GetById")

	u, err := s.repo.GetById(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on GetById", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "book fetched")

	return u, nil
}

func (s Service) Search(ctx context.Context, sp *entity.BookSearchParams) (*entity.BookList, error) {
	s.log.InfoContext(ctx, "entered Search")

	u, err := s.repo.Search(ctx, sp)
	if err != nil {
		s.log.ErrorContext(ctx, "error on Search", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "books fetched")

	return u, nil
}
