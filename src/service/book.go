package services

import (
	dbEntity "example_app/entity/db"
	httpEntity "example_app/entity/http"
	repository "example_app/repository/db"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

type BookService struct {
	bookRepository repository.BookRepositoryInterface
}

func BookServiceHandler() *BookService {
	return &BookService{
		bookRepository: repository.BookRepositoryHandler(),
	}
}

type BookServiceInterface interface {
	GetBooksList(page int, count int) []httpEntity.BookResponse
	GetBookById(id int) *httpEntity.BookDetailResponse
	StoreBook(payload httpEntity.BookRequest) bool
	UpdateBookById(id int, payload httpEntity.BookRequest) bool
	DeleteBook(id int) *httpEntity.BookResponse
}

func (service *BookService) GetBooksList(page int, count int) []httpEntity.BookResponse {
	books, _ := service.bookRepository.GetBooksList(page, count)
	result := []httpEntity.BookResponse{}
	copier.Copy(&result, &books)
	return result
}

func (service *BookService) GetBookById(id int) *httpEntity.BookDetailResponse {
	book := &dbEntity.Book{}
	service.bookRepository.GetBookById(id, book)

	result := &httpEntity.BookDetailResponse{}
	if book != nil {
		copier.Copy(result, book)
	}
	return result
}

func (service *BookService) StoreBook(payload httpEntity.BookRequest) bool {
	now := time.Now()
	book := &dbEntity.Book{
		Title:         payload.Title,
		Synopsis:      payload.Synopsis,
		PublishedDate: payload.PublishedDate,
		AuthorId:      payload.AuthorId,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	err := service.bookRepository.StoreBook(book)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *BookService) UpdateBookById(id int, payload httpEntity.BookRequest) bool {
	now := time.Now()
	book := &dbEntity.Book{
		Title:         payload.Title,
		Synopsis:      payload.Synopsis,
		PublishedDate: payload.PublishedDate,
		AuthorId:      payload.AuthorId,
		UpdatedAt:     &now,
	}
	err := service.bookRepository.UpdateBookById(id, book)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *BookService) DeleteBook(id int) *httpEntity.BookResponse {
	book := dbEntity.Book{}
	result := service.bookRepository.DeleteBook(id, &book)

	output := &httpEntity.BookResponse{}
	if result == nil {
		copier.Copy(output, book)
	}
	return output
}
