package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/bookstore/api/handler/dto"
	"github.com/guneyin/bookstore/api/middleware"
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

// GetBookList
// @Summary Get book list.
// @Description List all the books.
// @Tags book
// @Accept json
// @Produce json
// @Success 200 {object} middleware.ResponseHTTP{data=dto.BookListResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /book/list [get]
func (h BookHandler) GetBookList(c *fiber.Ctx) error {
	obj, err := h.svc.GetList(c.Context())
	if err != nil {
		return err
	}

	data := dto.BookListFromEntity(obj)

	return middleware.OK(c, fmt.Sprintf("%d books fetched", len(*data)), data)
}

// GetBookById
// @Summary Get book.
// @Description Get book by id.
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.BookResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /book/{id} [get]
func (h BookHandler) GetBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("invalid book is '%s'", c.Params("id"))
	}

	obj, err := h.svc.GetById(c.Context(), uint(id))
	if err != nil {
		return err
	}

	data := dto.BookFromEntity(obj)

	return middleware.OK(c, "book fetched", data)
}

// SearchBook
// @Summary Search book.
// @Description Search book by given parameters.
// @Tags book
// @Accept json
// @Produce json
// @Param search body dto.BookSearchRequest true "Search book"
// @Success 200 {object} middleware.ResponseHTTP{data=dto.BookListResponse}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /book/search [post]
func (h BookHandler) SearchBook(c *fiber.Ctx) error {
	sp := new(dto.BookSearchRequest)

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
