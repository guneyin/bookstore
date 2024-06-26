package cart

import (
	"context"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/mail"
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

	s.sendOrderConfirmMail(ctx, order.ID)

	return order, nil
}

func (s Service) GetOrderById(ctx context.Context, id uint) (*entity.Order, error) {
	s.log.InfoContext(ctx, "entered GetOrderById")

	order, err := s.repo.GetOrder(ctx, id)
	if err != nil {
		s.log.ErrorContext(ctx, "error on GetOrderById", slog.String("msg", err.Error()))

		return nil, err
	}

	s.log.InfoContext(ctx, "order fetched successfully")

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

func (s Service) sendOrderConfirmMail(ctx context.Context, orderId uint) {
	or, err := cart.GetOrderResult(ctx, orderId)
	if err != nil {
		s.log.ErrorContext(ctx, "error on sendOrderConfirmMail", slog.String("msg", err.Error()))

		return
	}

	if len(or) == 0 {
		s.log.WarnContext(ctx, "order not found")
	}

	order := or[0]
	mail.NewComposer(ctx).
		Name(order.UserName).
		To(order.UserEmail).
		Subject("Siparişiniz başarıyla alınmıştır").
		Template(mail.OrderConfirmationTemplate).
		Data(or).
		Send()

	s.log.InfoContext(ctx, "mail sent")
}
