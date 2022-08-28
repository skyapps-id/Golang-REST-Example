package dto

import "time"

type (
	BookRequest struct {
		Title string `json:"title" binding:"required"`
	}

	BookResponses struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	BookResponse struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
