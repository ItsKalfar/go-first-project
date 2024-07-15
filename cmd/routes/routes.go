package routes

import (
	"firstproject/cmd/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux();
	
	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("POST /createUser", handlers.CreateUsers)
	return mux
}