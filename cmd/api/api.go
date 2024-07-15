package api

import (
	"firstproject/cmd/middleware"
	"firstproject/cmd/routes"
	"net/http"
)



type APIServer struct{
	addr string
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
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
	return server.ListenAndServe()
}
