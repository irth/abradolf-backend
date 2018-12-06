package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irth/abradolf-backend/internal/abradolf"
)

func main() {
	a := abradolf.Abradolf{}
	m := mux.NewRouter()
	a.RegisterHandlers(m)
	http.ListenAndServe(":8080", m)
}
