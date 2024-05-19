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

	sc, err := s.repo.AddToCart(ctx, uId, bId, qty)
	if err != nil {
		s.log.ErrorContext(ctx, "error on AddToCart", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "book added to cart successfully")

	return sc, nil
}

func (s Service) GetCart(ctx context.Context, id uint) ([]entity.CartResult, error) {
	s.log.InfoContext(ctx, "entered GetCart")

	sc, err := s.repo.GetCart(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on GetCart", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "cart fetched successfully")

	return sc, nil
}

func (s Service) PlaceOrder(ctx context.Context, id uint) (*entity.Order, error) {
	s.log.InfoContext(ctx, "entered PlaceOrder")

	order, err := s.repo.PlaceOrder(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on PlaceOrder", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "order created successfully")

	return order, nil
}

func (s Service) GetOrdersByUserId(ctx context.Context, id uint) ([]entity.Order, error) {
	s.log.InfoContext(ctx, "entered GetOrdersByUserId")

	orders, err := s.repo.GetOrdersByUserId(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on GetOrdersByUserId", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "orders fetched successfully")

	return orders, nil
}
