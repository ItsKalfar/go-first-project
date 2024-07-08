package api

import (
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
	return http.ListenAndServe(s.addr, router)
}