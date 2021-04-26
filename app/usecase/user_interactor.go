package usecase

import (
	"app/app/domain"
	"fmt"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Remove(ID int) error {
	err := interactor.UserRepository.DeleteById(ID)
	return err
}

func (interactor *UserInteractor) UpdateUser(user domain.User) (domain.User, error) {
	user, err := interactor.UserRepository.Update(user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("UpdateUser error")
		return domain.User{}, err
	}
	return user, nil
}
