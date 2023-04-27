package users

import (
	eu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/users"
	ru "github.com/berrylradianh/go_berryl-radian-hamesha/modules/repository/users"
)

type Usecase struct {
	Repo ru.Repository
}

func (usecase Usecase) GetAllUsers() ([]eu.User, error) {
	users, err := usecase.Repo.GetAllUsers()
	return users, err
}

func (usecase Usecase) GetUser(id int) (eu.User, error) {
	user, err := usecase.Repo.GetUser(id)
	return user, err
}

func (usecase Usecase) CreateUser(user eu.User) error {
	err := usecase.Repo.CreateUser(user)
	return err
}

func (usecase Usecase) UpdateUser(id int, user eu.User) error {
	err := usecase.Repo.UpdateUser(id, user)
	return err
}

func (usecase Usecase) DeleteUser(id int) error {
	err := usecase.Repo.DeleteUser(id)
	return err
}

func (usecase Usecase) SearchUser(id int) error {
	err := usecase.Repo.SearchUser(id)
	return err
}
