package abradolf

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AuthResource struct{}

func (a *AuthResource) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/login", a.Login).Methods("POST")
	r.HandleFunc("/logout", a.Logout).Methods("POST")
	r.HandleFunc("/register", a.Register).Methods("POST")
}

func (a *AuthResource) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: implement AuthResource.Login
	w.Write([]byte("login"))
}

func (a *AuthResource) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: implement AuthResource.Logout
	w.Write([]byte("logout"))
}

func (a *AuthResource) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: implement AuthResource.Register
	w.Write([]byte("register"))
}
