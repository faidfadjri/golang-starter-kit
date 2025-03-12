package usecase

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/models"
	"akastra-mobile-api/src/infrastructure/repositories"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user entities.User) error
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

// @constructor
func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *userUsecase) GetUserByID(id int) (*models.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecase) CreateUser(user entities.User) error {
	return u.userRepo.CreateUser(user)
}
