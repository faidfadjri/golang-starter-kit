package usecase

import (
	"akastra-mobile-api/src/app/entities"
	"akastra-mobile-api/src/infrastructure/repositories"
	"regexp"
	"strings"
)

type AuthUsecase interface {
	Register(entities.UserRegisterPayload) (entities.UserRegisterPayload, error)
}

type authUseCase struct {
	authRepo repositories.AuthRepository
}

func NewAuthUsecase(authRepo repositories.AuthRepository) AuthUsecase {
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

func generateUsername(fullname string) string {
	fullname = strings.ToLower(fullname)

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	username := reg.ReplaceAllString(fullname, "")

	if len(username) > 15 {
		username = username[:15]
	}

	return username
}
