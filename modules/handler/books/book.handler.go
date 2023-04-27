package books

import (
	"fmt"
	"net/http"
	"strconv"

	eb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/books"
	ub "github.com/berrylradianh/go_berryl-radian-hamesha/modules/usecase/books"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase ub.Usecase
}

func (handler Handler) GetAllBooks() echo.HandlerFunc {
	return func(e echo.Context) error {
		var books []eb.Book

		books, err := handler.Usecase.GetAllBooks()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all books",
			"books":   books,
		})
	}
}

func (handler Handler) GetBook() echo.HandlerFunc {
	return func(e echo.Context) error {
		var book eb.Book
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		book, err = handler.Usecase.GetBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get book",
			"book":    book,
		})
	}
}

func (handler Handler) CreateBook() echo.HandlerFunc {
	var book eb.Book

	return func(e echo.Context) error {
		if err := e.Bind(&book); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		if err := e.Validate(book); err != nil {
			message := ""
			for _, e := range err.(validator.ValidationErrors) {
				message += fmt.Sprintf("%s is required ", e.Field())
			}
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
				"errors":  err.Error(),
			})
		}

		err := handler.Usecase.CreateBook(book)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new book",
			"book":    &book,
		})
	}
}

func (handler Handler) UpdateBook() echo.HandlerFunc {
	var book eb.Book

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&book); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.Usecase.UpdateBook(id, book)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update book",
		})
	}
}

func (handler Handler) DeleteBook() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = handler.Usecase.DeleteBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Book`",
		})
	}
}
