package bootstrap

import (
	"akastra-mobile-api/src/app/usecase"
	"akastra-mobile-api/src/infrastructure/database"
	"akastra-mobile-api/src/infrastructure/repositories"
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

	authRepo := repositories.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	return &Dependencies{
		AuthHandler: authHandler,
	}
}
