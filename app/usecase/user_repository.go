package usecase

import "app/app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindAll() (domain.Users, error)
	FindById(int) (domain.User, error)
	DeleteById(int) error
}
