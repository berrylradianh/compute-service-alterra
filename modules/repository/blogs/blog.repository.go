package blogs

import (
	ebl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/blogs"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllBlogs() ([]ebl.Blog, error) {
	var blogs []ebl.Blog
	result := repo.DB.Find(&blogs)

	return blogs, result.Error
}

func (repo Repository) GetBlog(id int) (ebl.Blog, error) {
	var blog ebl.Blog
	result := repo.DB.First(&blog, id)

	return blog, result.Error
}

func (repo Repository) CreateBlog(blog ebl.Blog) error {
	result := repo.DB.Create(&blog)
	return result.Error
}

func (repo Repository) UpdateBlog(id int, blog ebl.Blog) error {
	result := repo.DB.Model(&blog).Where("id = ?", id).Updates(&blog)
	return result.Error
}

func (repo Repository) DeleteBlog(id int) error {
	result := repo.DB.Delete(&ebl.Blog{}, id)
	return result.Error
}

func (repo Repository) SearchBlog(id int) error {
	result := repo.DB.First(&ebl.Blog{}, id)
	return result.Error
}
