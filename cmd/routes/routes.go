package routes

import (
	"firstproject/cmd/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux();
	mux.HandleFunc("GET /api/users", handlers.GetUsers);
	mux.HandleFunc("POST /api/about", handlers.CreateUsers);
	return mux
}