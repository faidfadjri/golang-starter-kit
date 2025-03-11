package routes

import (
	"database/sql"
	"net/http"

	"akastra-mobile-api/src/app/usecase"
	"akastra-mobile-api/src/infrastructure/repositories"
	"akastra-mobile-api/src/interface/handler"

	"github.com/go-chi/chi/v5"
)

func InitRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/users", UserRouter(userHandler))
	})

	return r
}
