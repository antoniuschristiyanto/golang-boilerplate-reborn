package controller

import (
	srv "example_app/service"

	"github.com/gin-gonic/gin"
)

func LoadRouter(routers *gin.Engine) {
	user := &UserRouterLoader{}
	book := &BookRouterLoader{}

	user.UserRouter(routers)
	book.BookRouter(routers)
}

type UserRouterLoader struct {
}

type BookRouterLoader struct {
}

func (rLoader *UserRouterLoader) UserRouter(router *gin.Engine) {
	handler := &UserController{
		UserService: srv.UserServiceHandler(),
	}
	rLoader.routerDefinition(router, handler)
}

func (rLoader *BookRouterLoader) BookRouter(router *gin.Engine) {
	handler := &BookController{
		BookService: srv.BookServiceHandler(),
	}
	rLoader.routerDefinition(router, handler)
}

func (rLoader *UserRouterLoader) routerDefinition(router *gin.Engine, handler *UserController) {
	group := router.Group("v1/users")
	group.GET("", handler.GetUsers)
	group.GET(":id", handler.GetUserByID)
	group.PUT(":id", handler.UpdateUsersByID)
}

func (rLoader *BookRouterLoader) routerDefinition(router *gin.Engine, handler *BookController) {
	group := router.Group("v1/books")
	group.GET("", handler.GetBooksList)
	group.GET(":id", handler.GetBookById)
	group.POST("", handler.StoreBook)
	group.PUT(":id", handler.UpdateBookById)
	group.DELETE(":id", handler.DeleteBook)
}
