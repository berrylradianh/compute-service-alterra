package users

import (
	eu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/users"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllUsers() ([]eu.User, error) {
	var users []eu.User
	result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := repo.DB.Find(&users)

	return users, result.Error
}

func (repo Repository) GetUser(id int) (eu.User, error) {
	var users eu.User
	result := repo.DB.Preload("Blogs", "deleted_at IS NULL").First(&users, id)
	// result := repo.DB.First(&users, id)

	return users, result.Error
}

func (repo Repository) CreateUser(user eu.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}

func (repo Repository) UpdateUser(id int, user eu.User) error {
	result := repo.DB.Model(&user).Where("id = ?", id).Updates(&user)
	return result.Error
}

func (repo Repository) DeleteUser(id int) error {
	result := repo.DB.Delete(&eu.User{}, id)
	return result.Error
}

func (repo Repository) SearchUser(id int) error {
	result := repo.DB.First(&eu.User{}, id)
	return result.Error
}
