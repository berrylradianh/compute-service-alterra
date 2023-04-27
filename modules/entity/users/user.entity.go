package users

import (
	b "github.com/berrylradianh/go_berryl-radian-hamesha/modules/response/blogs"
	"gorm.io/gorm"
)

// func InitMigrate() {
// 	db.DB.AutoMigrate(User{})
// }

type User struct {
	*gorm.Model

	Name     string           `json:"name" form:"name" validate:"required"`
	Email    string           `json:"email" form:"email" validate:"required,email"`
	Password string           `json:"password" form:"password" validate:"required"`
	Blogs    []b.BlogResponse `json:"blogs" form:"blogs"`
}
