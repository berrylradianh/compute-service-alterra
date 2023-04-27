package books

import (
	"gorm.io/gorm"
)

// func InitMigrate() {
// 	db.DB.AutoMigrate(Book{})
// }

type Book struct {
	*gorm.Model

	Title     string `json:"title" form:"title" validate:"required"`
	Writer    string `json:"writer" form:"writer" validate:"required"`
	Publisher string `json:"publisher" form:"publisher" validate:"required"`
}
