package blogs

import (
	"fmt"
	"net/http"
	"strconv"

	ebl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/blogs"
	ubl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/usecase/blogs"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase ubl.Usecase
}

func (handler Handler) GetAllBlogs() echo.HandlerFunc {
	return func(e echo.Context) error {
		var blogs []ebl.Blog

		blogs, err := handler.Usecase.GetAllBlogs()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all blogs",
			"blogs":   blogs,
		})
	}
}

func (handler Handler) GetBlog() echo.HandlerFunc {
	return func(e echo.Context) error {
		var blog ebl.Blog
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchBlog(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		blog, err = handler.Usecase.GetBlog(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get blog",
			"blog":    blog,
		})
	}
}

func (handler Handler) CreateBlog() echo.HandlerFunc {
	var blog ebl.Blog

	return func(e echo.Context) error {
		if err := e.Bind(&blog); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		if err := e.Validate(blog); err != nil {
			message := ""
			for _, e := range err.(validator.ValidationErrors) {
				message += fmt.Sprintf("%s is required ", e.Field())
			}
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
				"errors":  err.Error(),
			})
		}

		err := handler.Usecase.CreateBlog(blog)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new blog",
			"blog":    &blog,
		})
	}
}

func (handler Handler) UpdateBlog() echo.HandlerFunc {
	var blog ebl.Blog

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchBlog(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&blog); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.Usecase.UpdateBlog(id, blog)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update blog",
		})
	}
}

func (handler Handler) DeleteBlog() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchBlog(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = handler.Usecase.DeleteBlog(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Blog",
		})
	}
}
