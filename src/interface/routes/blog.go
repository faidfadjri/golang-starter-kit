package routes

import (
	"akastra-mobile-api/src/interface/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func BlogRouter(b *handler.BlogHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", b.GetArticles)    
	return r
}
