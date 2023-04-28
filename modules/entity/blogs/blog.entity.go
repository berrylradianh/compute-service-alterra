package blogs

import (
	"gorm.io/gorm"
)

type Blog struct {
	*gorm.Model

	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	UserID  int    `json:"user_id" form:"user_id" validate:"required"`
}
