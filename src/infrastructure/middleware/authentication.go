package middleware

import (
	"akastra-mobile-api/src/infrastructure/jwt"
	"akastra-mobile-api/src/interface/response"
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware memvalidasi token JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			response.JSONResponse(w, http.StatusUnauthorized, response.NewErrorResponse("Unauthorized: Missing token", nil))
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, err := jwt.ValidateToken(tokenString)
		if err != nil {
			log.Println("ERROR: Invalid token:", err)
			response.JSONResponse(w, http.StatusUnauthorized, response.NewErrorResponse("Unauthorized: Invalid token", nil))
			return
		}

		next.ServeHTTP(w, r)
	})
}
