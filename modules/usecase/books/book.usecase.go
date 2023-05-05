package books

import (
	eb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/books"
	rb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/repository/books"
)

type Usecase struct {
	Repo rb.Repository
}

func (usecase Usecase) GetAllBooks() ([]eb.Book, error) {
	books, err := usecase.Repo.GetAllBooks()
	return books, err
}

func (usecase Usecase) GetBook(id int) (eb.Book, error) {
	book, err := usecase.Repo.GetBook(id)
	return book, err
}

func (usecase Usecase) CreateBook(book eb.Book) error {
	err := usecase.Repo.CreateBook(book)
	return err
}

func (usecase Usecase) UpdateBook(id int, book eb.Book) error {
	err := usecase.Repo.UpdateBook(id, book)
	return err
}

func (usecase Usecase) DeleteBook(id int) error {
	err := usecase.Repo.DeleteBook(id)
	return err
}

func (usecase Usecase) SearchBook(id int) error {
	err := usecase.Repo.SearchBook(id)
	return err
}
