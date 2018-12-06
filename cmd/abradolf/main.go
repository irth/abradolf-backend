package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irth/abradolf-backend/internal/abradolf"
)

func main() {
	a := abradolf.Abradolf{}
	r := mux.NewRouter()
	r.StrictSlash(true)
	a.RegisterHandlers(r)
	http.ListenAndServe(":8080", r)
}
