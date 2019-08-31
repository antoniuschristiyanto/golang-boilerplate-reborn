package db

import "time"

type Book struct {
	Id            uint       `gorm:"primary_key" json:"id"`
	Title         string     `gorm:"column:title" json:"title"`
	Synopsis      string     `gorm:"column:synopsis" json:"synopsis"`
	PublishedDate *time.Time `gorm:"column:published_date" json:"published_date"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	User          *User      `gorm:"auto_preload, foreignkey:UserId, association_foreignkey:ID"`
	AuthorId      uint       `gorm:"column:author_id"`
}

func (Book) TableName() string {
	return "books"
}
