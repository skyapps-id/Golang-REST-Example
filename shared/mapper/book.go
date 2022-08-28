package mapper

import (
	"golang-rest-example/model"
	"golang-rest-example/shared/dto"
)

func MapBookToBookResponses(data []model.Book) []dto.BookResponses {
	var books []dto.BookResponses

	for _, row := range data {
		books = append(books, dto.BookResponses{
			ID:    row.ID,
			Title: row.Title,
		})
	}
	return books
}

func MapBookToBookResponse(data model.Book) *dto.BookResponse {
	return &dto.BookResponse{
		ID:        data.ID,
		Title:     data.Title,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
