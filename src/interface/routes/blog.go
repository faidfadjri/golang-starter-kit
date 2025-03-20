package routes

import (
	"akastra-mobile-api/src/interface/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func BlogRouter(a *handler.AuthHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/blog", a.Register)    
	return r
}
