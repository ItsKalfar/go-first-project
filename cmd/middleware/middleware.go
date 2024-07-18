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
		token := r.Header.Get("Authorization")

		if token == ""{
			utils.SendResponse(w, http.StatusUnauthorized, false, "Missing authentication header", nil)
			return
		}

		// token = token[len("Bearer "):] - To remove Bearer prefix

		err := utils.VerifyToken(token); 
		
		if err != nil {
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

