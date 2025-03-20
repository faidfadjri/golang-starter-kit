package routes

import (
	"akastra-mobile-api/src/app/bootstrap"
	"akastra-mobile-api/src/infrastructure/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRouter(deps *bootstrap.Dependencies) http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(api chi.Router) {
		r.Use(middleware.RateLimiter())
		api.Mount("/auth", AuthRouter(deps.AuthHandler))
		api.With(middleware.AuthMiddleware).Mount("/blog", BlogRouter(deps.BlogHandler))
	})

	return r
}
