package main

import (
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"github.com/irth/abradolf-backend/internal/db"

	"github.com/gorilla/mux"
	"github.com/irth/abradolf-backend/internal/app"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	db := db.Open("sqlite3", ":memory:")

	a := app.New(db)
	a.RegisterHandlers(r)

	fmt.Println("listening...")
	http.ListenAndServe(":8080", r)
}
