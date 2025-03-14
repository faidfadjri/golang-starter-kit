package repositories

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/database/models/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error)
	Login(user entities.UserCredentials) (users.User, error)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.UserRegisterPayload{}, err	
	}

	user.Password = string(hashedPassword)

	if err := r.db.Create(&user).Error; err != nil {
		return entities.UserRegisterPayload{}, err
	}
	return user, nil
}

func (r *authRepository) Login(user entities.UserCredentials) (users.User, error) {	
	var userDB users.User

	result := r.db.Where("username = ?", user.EmailOrUsername).
		Or("email = ?", user.EmailOrUsername).
		First(&userDB)

	if result.Error != nil {
		return users.User{}, fmt.Errorf("user not found")
	} 
	fmt.Println("User email:", userDB.Email)
	return userDB, nil
}
