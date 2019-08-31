package db

import (
	"errors"
	dbEntity "example_app/entity/db"
	connection "example_app/util/helper/mysqlconnection"

	"github.com/jinzhu/gorm"
)

type BookRepository struct {
	DB gorm.DB
}

func BookRepositoryHandler() *BookRepository {
	return &BookRepository{DB: *connection.GetConnection()}
}

type BookRepositoryInterface interface {
	GetBooksList(limit int, offset int) ([]dbEntity.Book, error)
	GetBookById(id int, bookData *dbEntity.Book) error
	StoreBook(bookData *dbEntity.Book) error
	UpdateBookById(id int, bookData *dbEntity.Book) error
	DeleteBook(id int, bookData *dbEntity.Book) error
}

func (repository *BookRepository) GetBooksList(limit int, offset int) ([]dbEntity.Book, error) {
	books := []dbEntity.Book{}
	query := repository.DB.Table("books")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&books)
	return books, query.Error
}

func (repository *BookRepository) GetBookById(id int, bookData *dbEntity.Book) error {
	query := repository.DB.Preload("User")
	query = query.Where("id=?", id)
	query = query.First(bookData)
	return query.Error
}

func (repository *BookRepository) StoreBook(bookData *dbEntity.Book) error {
	query := repository.DB.Table("books")
	query = query.Create(bookData)
	return query.Error
}

func (repository *BookRepository) UpdateBookById(id int, bookData *dbEntity.Book) error {
	query := repository.DB.Table("books")
	query = query.Where("id=?", id)
	success := query.Updates(bookData).RowsAffected
	if success < 1 {
		return errors.New("No data affected")
	}
	return query.Error
}

func (repository *BookRepository) DeleteBook(id int, bookData *dbEntity.Book) error {
	book := &dbEntity.Book{}
	query := repository.DB.Table("books")
	query = query.Where("id=?", id)
	query = query.First(bookData)
	query = query.Delete(book)
	return query.Error
}
