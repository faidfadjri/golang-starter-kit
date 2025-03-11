package route

import (
	"akastra-mobile-api/src/interface/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllUsers)        // GET /api/v1/users
	r.Get("/{id}", h.GetUserByID)     // GET /api/v1/users/{id}
	r.Post("/", h.CreateUser)        // POST /api/v1/users

	return r
}
