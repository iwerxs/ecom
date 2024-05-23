package api

import (
	"database/sql"
	"ecom/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

// a new method to create an API server
func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

// initialize a router, register all the routes so register all the services and their dependencies
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	// Add a new Service here

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}