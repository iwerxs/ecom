package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {

}

// create a new constructor for the Handler
func NewHandler() *Handler{
	return &Handler{}
}

// Register Route function
func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleLogin).Methods("POST")
}

// a new method for the Handler
func (h *Handler) handleLogin(w http.ResponseWriter, r * http.Request){

}
func (h *Handler) handleRegister(w http.ResponseWriter, r * http.Request){

}