package services

import (
	"fmt"
	"testing"
)
import (
	modelDB "example_app/entity/db"
)
import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// SETUP
type repositoryBookMock struct {
	mock.Mock
}

func (repository *repositoryBookMock) GetBookById(id int, book *modelDB.Book) error {
	repository.Called(id, book)
	book.Id = uint(id)
	book.Title = "Indonesia"
	book.Synopsis = "74 Tahun Merdeka"
	return nil
}

func (repository *repositoryBookMock) DeleteBook(id int, book *modelDB.Book) error {
	repository.Called(id, book)
	book.Id = uint(id)
	book.Title = fmt.Sprintf("Removed - %s", book.Title)
	book.Synopsis = fmt.Sprintf("Removed - %s", book.Synopsis)
	return nil
}

func (repository *repositoryBookMock) StoreBook(book *modelDB.Book) error {
	repository.Called(book)
	book.Title = fmt.Sprintf("Stored - %s", book.Title)
	book.Synopsis = fmt.Sprintf("Stored - %s", book.Synopsis)
	return nil
}

func (repository *repositoryBookMock) UpdateBookById(id int, book *modelDB.Book) error {
	repository.Called(id, book)
	book.Id = uint(id)
	book.Title = fmt.Sprintf("Updated - %s", book.Title)
	book.Synopsis = fmt.Sprintf("Updated - %s", book.Synopsis)
	return nil
}

func (repository *repositoryBookMock) GetBooksList(limit int, offset int) ([]modelDB.Book, error) {
	repository.Called(limit, offset)
	books := []modelDB.Book{}
	const (
		ID1 = iota + 1
		ID2
		ID3
	)
	books = append(books, modelDB.Book{
		Id:    uint(ID1),
		Title: "Lorem",
		Body:  "Ipsum No Lorem",
	})
	books = append(books, modelDB.Book{
		Id:    uint(ID2),
		Title: "Ipsum",
		Body:  "Lorem",
	})
	books = append(books, modelDB.Book{
		Id:    uint(ID3),
		Title: "No Lorem",
		Body:  "Ipsum",
	})
	return books, nil
}

// TEST
func TestbookServiceGetBookByIdMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryBookMock{}

	book := &modelDB.Book{}
	var testId int = 1
	dbMockData.On("GetBookById", testId, book).Return(nil)

	bookService := bookService{&dbMockData}
	resultFuncService := bookService.GetbookById(testId)
	assert.Equal(t, uint(testId), resultFuncService.Id, "It should be same ID")
	assert.Equal(t, "Indonesia", resultFuncService.Title, "It should be same Title")
	assert.Equal(t, "74 Tahun Merdeka", resultFuncService.Body, "It should be same Synopsis")
}

func TestbookServiceGetBooksListMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryBookMock{}
	limit := 1
	offset := 3
	dbMockData.On("GetBooksList", limit, offset).Return([]modelDB.Book{}, nil)
	bookService := bookService{&dbMockData}
	resultFuncService := bookService.GetBooksList(limit, offset)
	assert.Equal(t, len(resultFuncService), 3, "It should be same length as Mock Data")
	assert.Equal(t, resultFuncService[0].Title, "Lorem", "It should be same NAME as Mock Data")
	assert.Equal(t, resultFuncService[1].Title, "Ipsum", "It should be same NAME as Mock Data")
	assert.Equal(t, resultFuncService[2].Title, "No Lorem", "It should be same NAME as Mock Data")
}
