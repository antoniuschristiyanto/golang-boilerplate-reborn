package http

import "time"

type BookRequest struct {
	Title         string     `json:"title"`
	Synopsis      string     `json:"synopsis"`
	PublishedDate *time.Time `json:"published_date"`
	AuthorId      uint       `json:"authod_id"`
}

type BookDetailResponse struct {
	Id            uint       `json:"id"`
	Title         string     `json:"title"`
	Synopsis      string     `json:"synopsis"`
	PublishedDate *time.Time `json:"published_date"`
	AuthorId      uint       `json:"authod_id"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type BookResponse struct {
	Id            uint       `json:"id"`
	Title         string     `json:"title"`
	PublishedDate *time.Time `json:"published_date"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
