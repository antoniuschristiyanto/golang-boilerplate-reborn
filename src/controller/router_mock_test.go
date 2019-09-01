package controller

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func LoadRouterTestMock() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	context, routers := gin.CreateTestContext(resp)

	routerLoader := &UserRouterLoader{}
	routerLoader.UserRouterTestMock(routers)
	routerBook := &BookRouterLoader{}
	routerBook.BookRouterTestMock(routers)

	return context, routers, resp
}

func (rLoader *UserRouterLoader) UserRouterTestMock(router *gin.Engine) {
	handler := &UserController{
		UserService: &UserServiceMock{},
	}
	rLoader.routerDefinition(router, handler)
}

func (rLoader *BookRouterLoader) BookRouterTestMock(router *gin.Engine) {
	handler := &BookController{
		BookService: &BookServiceMock{},
	}
	rLoader.routerDefinition(router, handler)
}
