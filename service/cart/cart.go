package cart

import (
	"context"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/repo/cart"
	"log/slog"
)

type Service struct {
	repo *cart.Repo
	log  *slog.Logger
}

func New(log *slog.Logger) *Service {
	return &Service{
		repo: cart.NewRepo(),
		log:  log,
	}
}

func (s Service) AddToCart(ctx context.Context, uId, bId, qty uint) ([]entity.CartResult, error) {
	s.log.InfoContext(ctx, "entered AddToCart")

	cart, err := s.repo.AddToCart(ctx, uId, bId, qty)
	if err != nil {
		s.log.ErrorContext(ctx, "error on AddToCart", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "book added to cart successfully")

	return cart, nil
}

func (s Service) GetChart(ctx context.Context, id uint) ([]entity.CartResult, error) {
	s.log.InfoContext(ctx, "entered GetChart")

	cart, err := s.repo.GetCart(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on GetChart", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "chart fetched successfully")

	return cart, nil
}
