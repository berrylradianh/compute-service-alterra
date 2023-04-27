package blogs

import (
	"gorm.io/gorm"
)

// func InitMigrate() {
// 	db.DB.AutoMigrate(Blog{})
// }

type Blog struct {
	*gorm.Model

	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	UserID  int    `json:"user_id" form:"user_id" validate:"required"`
	// User    u.UserResponse `json:"user" form:"user"`
}
