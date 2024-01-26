package controller

import (
	"fmt"
	"go_fiber/api"
	service "go_fiber/service/book"

	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	Service service.IBookService
}

type IBookController interface {
	Execute(c *fiber.Ctx) error
}

func NewBookController(bookService service.IBookService) IBookController {
	return &bookController{
		Service: bookService,
	}
}

func (b *bookController) Execute(c *fiber.Ctx) error {
	book, err := b.Service.FindAll()
	if err != nil {
		fmt.Println(err)
		return c.JSON(err)
	}

	fmt.Println(book)

	var apiRes api.BookResponse
	apiRes.Code = 200
	apiRes.Message = "Success"
	apiRes.Data = struct {
		Id   string "json:\"id\""
		Name string "json:\"name\""
	}{
		Id:   "1",
		Name: "Book 1",
	}

	return c.JSON(apiRes)
}
