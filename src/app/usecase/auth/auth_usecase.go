package auth

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/database/models/users"
	auth "akastra-mobile-api/src/infrastructure/repositories/auth"
	"regexp"
	"strings"
)

type authUseCase struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(authRepo auth.AuthRepository) AuthUsecase {
	return &authUseCase{authRepo: authRepo}
}

func (r *authUseCase) Register(user entities.UserRegisterPayload) (entities.UserRegisterPayload, error) {

	username := generateUsername(user.Fullname)
	user.Username = &username

	role_id := 2
	user.RoleId = &role_id

	user, err := r.authRepo.Register(user)
	if err != nil {
		return entities.UserRegisterPayload{}, err
	}
	return user, nil
}

func (r *authUseCase) Login(user entities.UserCredentials) (users.User, error) {
	validatedUser, err := r.authRepo.Login(user)
	if err != nil {
		return users.User{}, err
	}
	return validatedUser, nil
}

func generateUsername(fullname string) string {
	fullname = strings.ToLower(fullname)

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	username := reg.ReplaceAllString(fullname, "")

	if len(username) > 15 {
		username = username[:15]
	}

	return username
}
