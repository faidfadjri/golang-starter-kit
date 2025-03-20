package auth

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/database/models/users"
)

type AuthRepository interface {
	Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error)
	Login(user entities.UserCredentials) (users.User, error)
}