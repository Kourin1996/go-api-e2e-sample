package controller

import (
	"net/http"

	"github.com/Kourin1996/go-api-e2e-sample/domain/books"
	"github.com/labstack/echo/v4"
)

const BASE_PATH = "/books"

type BooksController struct {
	e        *echo.Echo
	bookRepo books.IBooksRepository
}

func NewController(
	e *echo.Echo,
	bookRepo books.IBooksRepository,
) *BooksController {
	controller := &BooksController{
		e:        e,
		bookRepo: bookRepo,
	}

	group := e.Group(BASE_PATH)
	group.POST("", controller.PostBook)
	group.GET("/:id", controller.GetBook)

	return controller
}

func (c *BooksController) PostBook(ctx echo.Context) error {
	dto := &books.CreateBookDTO{}

	if err := ctx.Bind(dto); err != nil {
		return err
	}

	res, err := c.bookRepo.CreateBook(dto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *BooksController) GetBook(ctx echo.Context) error {
	dto := &books.GetBookDTO{}

	if err := ctx.Bind(dto); err != nil {
		return err
	}

	res, err := c.bookRepo.GetBook(dto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
