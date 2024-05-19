package dto

import "github.com/guneyin/bookstore/entity"

type BookResponse struct {
	Id              uint    `json:"id"`
	Title           string  `json:"title"`
	Author          string  `json:"author"`
	PublicationYear uint    `json:"publicationYear"`
	Genre           string  `json:"genre"`
	Description     string  `json:"description"`
	CoverImage      string  `json:"coverImage"`
	Price           float64 `json:"price"`
}

type BookListResponse []BookResponse

type BookSearchRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func BookFromEntity(b *entity.Book) *BookResponse {
	return &BookResponse{
		Id:              b.ID,
		Title:           b.Title,
		Author:          b.Author,
		PublicationYear: b.PublicationYear,
		Genre:           b.Genre,
		Description:     b.Description,
		CoverImage:      b.CoverImage,
		Price:           b.Price,
	}
}

func BookListFromEntity(bl *entity.BookList) *BookListResponse {
	var list BookListResponse

	for _, item := range *bl {
		u := BookFromEntity(&item)
		list = append(list, *u)
	}

	return &list
}
