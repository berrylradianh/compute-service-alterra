package books

import (
	eb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/books"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllBooks() ([]eb.Book, error) {
	var books []eb.Book
	result := repo.DB.Find(&books)

	return books, result.Error
}

func (repo Repository) GetBook(id int) (eb.Book, error) {
	var book eb.Book
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").First(&users, id)
	result := repo.DB.First(&book, id)

	return book, result.Error
}

func (repo Repository) CreateBook(book eb.Book) error {
	result := repo.DB.Create(&book)
	// result := repo.DB.Create(&user).Preload("Blogs", "deleted_at IS NULL").First(&user, user.ID)
	return result.Error
}

func (repo Repository) UpdateBook(id int, book eb.Book) error {
	result := repo.DB.Model(&book).Where("id = ?", id).Updates(&book)
	return result.Error
}

func (repo Repository) DeleteBook(id int) error {
	result := repo.DB.Delete(&eb.Book{}, id)
	return result.Error
}

func (repo Repository) SearchBook(id int) error {
	result := repo.DB.First(&eb.Book{}, id)
	return result.Error
}
