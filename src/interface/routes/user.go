package routes

import (
	"akastra-mobile-api/src/interface/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllUsers)       
	r.Get("/{id}", h.GetUserByID)   
	r.Post("/", h.CreateUser)    

	return r
}
