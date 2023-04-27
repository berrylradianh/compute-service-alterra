package blogs

import (
	ebl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/blogs"
	rbl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/repository/blogs"
)

type Usecase struct {
	Repo rbl.Repository
}

func (usecase Usecase) GetAllBlogs() ([]ebl.Blog, error) {
	blogs, err := usecase.Repo.GetAllBlogs()
	return blogs, err
}

func (usecase Usecase) GetBlog(id int) (ebl.Blog, error) {
	blog, err := usecase.Repo.GetBlog(id)
	return blog, err
}

func (usecase Usecase) CreateBlog(blog ebl.Blog) error {
	err := usecase.Repo.CreateBlog(blog)
	return err
}

func (usecase Usecase) UpdateBlog(id int, blog ebl.Blog) error {
	err := usecase.Repo.UpdateBlog(id, blog)
	return err
}

func (usecase Usecase) DeleteBlog(id int) error {
	err := usecase.Repo.DeleteBlog(id)
	return err
}

func (usecase Usecase) SearchBlog(id int) error {
	err := usecase.Repo.SearchBlog(id)
	return err
}
