package middleware

import (
	"akastra-mobile-api/src/interface/response"
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

func RateLimiter() func(http.Handler) http.Handler {
	limiter := httprate.LimitByIP(60, time.Minute)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rec := &responseInterceptor{ResponseWriter: w, statusCode: http.StatusOK, bodyWritten: false}

			limiter(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if rec.statusCode != http.StatusTooManyRequests {
					next.ServeHTTP(rec, r)
				}
			})).ServeHTTP(rec, r)

			if rec.statusCode == http.StatusTooManyRequests {
				response.JSONResponse(w, http.StatusTooManyRequests, response.NewErrorResponse("Too many requests. Please try again later.", nil))
			}
		})
	}
}

type responseInterceptor struct {
	http.ResponseWriter
	statusCode  int
	bodyWritten bool
}

func (rec *responseInterceptor) WriteHeader(code int) {
	rec.statusCode = code
	if code == http.StatusTooManyRequests {
		rec.bodyWritten = true
		return
	}
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *responseInterceptor) Write(b []byte) (int, error) {
	if rec.statusCode == http.StatusTooManyRequests {
		rec.bodyWritten = true
		return len(b), nil
	}
	return rec.ResponseWriter.Write(b)
}
