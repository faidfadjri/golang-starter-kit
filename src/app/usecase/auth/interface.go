package auth

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/database/models/users"
)

type AuthUsecase interface {
	Register(entities.UserRegisterPayload) (entities.UserRegisterPayload, error)
	Login(entities.UserCredentials) (users.User, error)
}