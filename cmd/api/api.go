package api

import (
	"database/sql"
	"firstproject/cmd/middleware"
	"firstproject/cmd/routes"
	"log"
	"net/http"
)



type APIServer struct{
	addr string
	db *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := routes.SetupRoutes();

	Middleware := middleware.MiddlewareChain(
		middleware.RequestLogger, 
		middleware.AuthMiddleware,
	)

	server := http.Server{
		Addr: s.addr,
		Handler: Middleware(router),
	}

	log.Print("Server started on", s.addr);

	return server.ListenAndServe()
}
