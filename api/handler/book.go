package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/service/book"
	"log/slog"
)

const bookHandlerName = "book"

type BookHandler struct {
	svc *book.Service
}

var _ IHandler = (*UserHandler)(nil)

func NewBook(log *slog.Logger) IHandler {
	svc := book.New(log)

	return &BookHandler{svc}
}

func (h BookHandler) Name() string {
	return bookHandlerName
}

func (h BookHandler) SetRoutes(r fiber.Router) IHandler {
	g := r.Group(h.Name())
	g.Get("list", h.GetBookList)
	g.Get("/:id", h.GetBookById)
	g.Post("/search", h.SearchBook)

	return h
}

func (h BookHandler) GetBookList(c *fiber.Ctx) error {
	list, err := h.svc.GetList(c.Context())
	if err != nil {
		return err
	}

	data := dto.BookListFromEntity(list)

	return middleware.OK(c, fmt.Sprintf("%d books fetched", len(*data)), data)
}

func (h BookHandler) GetBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("invalid book is '%s'", c.Params("id"))
	}

	data, err := h.svc.GetById(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return middleware.OK(c, "book fetched", data)
}

func (h BookHandler) SearchBook(c *fiber.Ctx) error {
	sp := new(entity.BookSearchParams)

	err := c.BodyParser(sp)
	if err != nil {
		return err
	}

	list, err := h.svc.Search(c.Context(), sp)
	if err != nil {
		return err
	}

	data := dto.BookListFromEntity(list)

	return middleware.OK(c, fmt.Sprintf("%d books fetched", len(*data)), data)
}
