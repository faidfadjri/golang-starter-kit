package bootstrap

import (
	"akastra-mobile-api/src/app/usecase"
	"akastra-mobile-api/src/infrastructure/database"
	"akastra-mobile-api/src/infrastructure/repositories"
	"akastra-mobile-api/src/interface/handler"
	"log"
)

type Dependencies struct {
	UserHandler *handler.UserHandler
}

func InitDependencies() *Dependencies {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	return &Dependencies{
		UserHandler: userHandler,
	}
}
