package bootstrap

import (
	auth "akastra-mobile-api/src/app/usecase/auth"
	"akastra-mobile-api/src/infrastructure/database"
	repositoriesauth "akastra-mobile-api/src/infrastructure/repositories/auth"
	"akastra-mobile-api/src/interface/handler"
	"log"
)

type Dependencies struct {
	AuthHandler *handler.AuthHandler
}

func InitDependencies() *Dependencies {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	authRepo := repositoriesauth.NewAuthRepository(db)
	authUsecase := auth.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	return &Dependencies{
		AuthHandler: authHandler,
	}
}
