package gen

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/entity"
	"github.com/guneyin/bookstore/repo/book"
	"gorm.io/gorm"
	"log/slog"
	"math/rand"
	"strings"
)

const urlBooks = "https://freetestapi.com/api/v1/books"

type BookDTO struct {
	Title           string   `json:"title"`
	Author          string   `json:"author"`
	PublicationYear uint     `json:"publication_year"`
	Genre           []string `json:"genre"`
	Description     string   `json:"description"`
	CoverImage      string   `json:"cover_image"`
}

func generateBookData(ctx context.Context, r *resty.Request, db *gorm.DB) error {
	log.InfoContext(ctx, "generating book data..")

	ul, err := fetchData[BookDTO](ctx, r, db, urlBooks)
	if err != nil {
		return err
	}

	truncateTable(db, &entity.Book{})

	repo := book.NewRepo()
	for _, item := range ul {
		err = repo.Create(ctx, &entity.Book{
			Title:           item.Title,
			Author:          item.Author,
			PublicationYear: item.PublicationYear,
			Genre:           strings.Join(item.Genre, " "),
			Description:     item.Description,
			CoverImage:      item.CoverImage,
			Price:           (rand.Float64() * 50) + 50,
		})
		if err != nil {
			log.WarnContext(ctx, fmt.Sprintf("%s could not created", item.Title), slog.String("err", err.Error()))
		}
	}

	return nil
}
