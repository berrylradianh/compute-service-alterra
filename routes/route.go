package routes

import (
	db "github.com/berrylradianh/go_berryl-radian-hamesha/databases"
	s "github.com/berrylradianh/go_berryl-radian-hamesha/modules/services"
	"github.com/go-playground/validator"

	hu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/handler/users"
	ru "github.com/berrylradianh/go_berryl-radian-hamesha/modules/repository/users"
	uu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/usecase/users"

	hb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/handler/books"
	rb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/repository/books"
	ub "github.com/berrylradianh/go_berryl-radian-hamesha/modules/usecase/books"

	hbl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/handler/blogs"
	rbl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/repository/blogs"
	ubl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/usecase/blogs"

	"github.com/labstack/echo/v4"
)

var (
	userRepo    ru.Repository
	userHandler hu.Handler
	userUsecase uu.Usecase

	bookRepo    rb.Repository
	bookUsecase ub.Usecase
	bookHandler hb.Handler

	blogRepo    rbl.Repository
	blogUsecase ubl.Usecase
	blogHandler hbl.Handler
)

func declare() {
	userRepo = ru.Repository{DB: db.DB}
	userUsecase = uu.Usecase{Repo: userRepo}
	userHandler = hu.Handler{Usecase: userUsecase}

	bookRepo = rb.Repository{DB: db.DB}
	bookUsecase = ub.Usecase{Repo: bookRepo}
	bookHandler = hb.Handler{Usecase: bookUsecase}

	blogRepo = rbl.Repository{DB: db.DB}
	blogUsecase = ubl.Usecase{Repo: blogRepo}
	blogHandler = hbl.Handler{Usecase: blogUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.Validator = &s.CustomValidator{Validator: validator.New()}
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	user := e.Group("/users")
	user.GET("", userHandler.GetAllUsers())
	user.GET("/:id", userHandler.GetUser())
	user.POST("", userHandler.CreateUser())
	user.DELETE("/:id", userHandler.DeleteUser())
	user.PUT("/:id", userHandler.UpdateUser())

	book := e.Group("/books")
	book.GET("", bookHandler.GetAllBooks())
	book.GET("/:id", bookHandler.GetBook())
	book.POST("", bookHandler.CreateBook())
	book.DELETE("/:id", bookHandler.DeleteBook())
	book.PUT("/:id", bookHandler.UpdateBook())

	blog := e.Group("/blogs")
	blog.GET("", blogHandler.GetAllBlogs())
	blog.GET("/:id", blogHandler.GetBlog())
	blog.POST("", blogHandler.CreateBlog())
	blog.DELETE("/:id", blogHandler.DeleteBlog())
	blog.PUT("/:id", blogHandler.UpdateBlog())

	return e
}
