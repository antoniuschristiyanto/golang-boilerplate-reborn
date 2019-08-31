package controller

import (
	httpEntity "example_app/entity/http"
	services "example_app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService services.BookServiceInterface
}

func (service *BookController) GetBooksList(context *gin.Context) {
	queryparam := Limitofset{}
	err := context.ShouldBindQuery(&queryparam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if queryparam.Limit == 0 {
		queryparam.Limit = 10
	}
	result := service.BookService.GetBooksList(queryparam.Limit, queryparam.Offset)
	context.JSON(http.StatusOK, result)
}

func (service *BookController) GetBookById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	result := service.BookService.GetBookById(id)
	if result == nil {
		context.JSON(http.StatusOK, gin.H{})
		return
	}
	context.JSON(http.StatusOK, result)
}

func (service *BookController) StoreBook(context *gin.Context) {
	payload := httpEntity.BookRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.BookService.StoreBook(payload)
	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})

}

func (service *BookController) UpdateBookById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	payload := httpEntity.BookRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.BookService.UpdateBookById(id, payload)

	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func (service *BookController) DeleteBook(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	result := service.BookService.DeleteBook(id)

	if result.Id == 0 {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusOK, result)
}
