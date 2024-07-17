package routes

import (
	"database/sql"
	"firstproject/cmd/handlers"
	"net/http"
)

func SetupRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux();

	categoriesHandler := handlers.CategoryHandler(db)
	categoriesHandler.CategoryRouter(mux)

	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("POST /createUser", handlers.CreateUsers)

	return mux
}