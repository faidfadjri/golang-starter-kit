package routes

import (
	"akastra-mobile-api/src/app/bootstrap"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRouter(deps *bootstrap.Dependencies) http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/auth", AuthRouter(deps.AuthHandler))
	})

	return r
}
