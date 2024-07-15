package middleware

import (
	"firstproject/cmd/utils"
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.HandlerFunc

func RequestLogger (next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		log.Printf(r.Method, r.URL.Path);
		next.ServeHTTP(w, r)
	}
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		// Check for authentication
		token := r.Header.Get("Authorization")

		log.Printf("Auth Middleware: Received token: %s", token)

		if token != "1234" {
			utils.SendResponse(w, http.StatusUnauthorized, false, "Unauthorized", nil)
			return
		}
		
		next.ServeHTTP(w, r);
	}
}

func MiddlewareChain (middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}

		return next.ServeHTTP
	}
}

