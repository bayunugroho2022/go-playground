package handlers

import (
	"go_fiber/controller"
	repositories "go_fiber/repositories/book"
	service "go_fiber/service/book"
)


func (h *MyHandler) BookHandler() controller.IBookController {
	bookRepositories := repositories.NewBookRepositories(h.Application.PostgresConnection)
	bookService := service.NewUserService(bookRepositories)
	bookController := controller.NewBookController(bookService)
	return bookController
}