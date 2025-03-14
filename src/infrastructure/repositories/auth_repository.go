package repositories

import (
	"akastra-mobile-api/src/app/entities"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error)
	login(user entities.UserLoginPayload) (bool, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error) {
	var existingUser entities.UserRegisterPayload
	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return entities.UserRegisterPayload{}, fmt.Errorf("email already exists")
	}

	if err := r.db.Create(&user).Error; err != nil {
		return entities.UserRegisterPayload{}, err
	}
	return user, nil
}

func (r *authRepository) login(user entities.UserLoginPayload) (bool, error) {
	var userDB entities.UserRegisterPayload
	if err := r.db.Where("email = ?", user.EmailOrUsername).First(&userDB).Error; err != nil {
		return false, err
	}
	return true, nil
}