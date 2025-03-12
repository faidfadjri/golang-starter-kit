package repositories

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/database/models/users"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]users.User, error)
	GetUserByID(id int) (*users.User, error)
	CreateUser(user entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]users.User, error) {
	var users []users.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id int) (*users.User, error) {
	var user users.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user entities.User) error {
	return r.db.Create(&user).Error
}
