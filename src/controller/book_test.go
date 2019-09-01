package controller

import (
	"encoding/json"
	httpEntity "example_app/entity/http"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type BookServiceMock struct{}

func (service *BookServiceMock) GetBookById(id int) *httpEntity.BookDetailResponse {
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	return &httpEntity.BookDetailResponse{
		Id:            uint(id),
		Title:         "Indonesia",
		Body:          "74 Tahun Merdeka",
		AuthorId:      35,
		PublishedDate: &t,
		CreatedAt:     &t,
		UpdatedAt:     &t,
	}
}

func (service *BookServiceMock) GetBooksList(page int, count int) []httpEntity.BookResponse {
	books := []httpEntity.BookResponse{}
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	const (
		ID1 = iota + 1
		ID2
		ID3
	)
	books = append(books, httpEntity.BookResponse{
		Id:        uint(ID1),
		Title:     "Lorem",
		CreatedAt: &t,
		UpdatedAt: &t,
	})
	books = append(books, httpEntity.BookResponse{
		Id:        uint(ID2),
		Title:     "Ipsum",
		CreatedAt: &t,
		UpdatedAt: &t,
	})
	books = append(books, httpEntity.BookResponse{
		Id:        uint(ID1),
		Title:     "No Lorem",
		CreatedAt: &t,
		UpdatedAt: &t,
	})
	return books
}

func (service *BookServiceMock) UpdateBookById(id int, payload httpEntity.BookRequest) bool {
	return true
}

func (service *BookServiceMock) DeleteBook(id int) *httpEntity.BookResponse {
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	return &httpEntity.BookResponse{
		Id:        uint(id),
		Title:     "Indonesia",
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}

func (service *BookServiceMock) StoreBook(payload httpEntity.BookRequest) bool {
	return true
}

func TestBookGetByIDMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest uint = 1
	url := "/v1/books/" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := httpEntity.BookDetailResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)

	assert.Equal(err, nil, "should have no error")
	assert.Equal(res.Id, idTest, "It should be same ID")
	assert.Equal(res.Title, "Indonesia", "It should be same Title")
	assert.Equal(res.Body, "74 Tahun Merdeka", "It should be same Synopsis")
}

func TestGetBookListMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	url := "/v1/books"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := []httpEntity.BookResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)

	assert.Equal(err, nil, "should have no error")
	assert.Equal(len(res) >= 0, true, "length must in minimum value")
	assert.Equal(len(res) == 3, true, "length value must match with mock data")
}
