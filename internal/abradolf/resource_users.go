package abradolf

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UsersResource struct{}

func (u *UsersResource) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("", u.List).Methods("GET")
	r.HandleFunc("/{id}", u.Get).Methods("GET", "DELETE")
}

func (u *UsersResource) List(w http.ResponseWriter, r *http.Request) {
	// TODO: implement UsersResource.List
	w.Write([]byte("list"))
}

func (u *UsersResource) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: implement UsersResource.Get
	w.Write([]byte("get"))
}
