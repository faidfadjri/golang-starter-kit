package routes

import (
	"akastra-mobile-api/src/interface/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(a *handler.AuthHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", a.Register)    

	return r
}
