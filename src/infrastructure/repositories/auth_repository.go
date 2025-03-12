package repositories

import (
	"akastra-mobile-api/src/app/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return entities.UserRegisterPayload{}, err
	}
	return user, nil
}