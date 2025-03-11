package usecase

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/repositories"
)

type UserUsecase interface {
	GetAllUsers() ([]entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	CreateUser(user entities.User) error
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetAllUsers() ([]entities.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *userUsecase) GetUserByID(id int) (*entities.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecase) CreateUser(user entities.User) error {
	return u.userRepo.CreateUser(user)
}
