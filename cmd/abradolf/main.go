package main

import (
	"fmt"
	"net/http"

	"github.com/irth/abradolf-backend/internal/db"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gorilla/mux"
	"github.com/irth/abradolf-backend/internal/app"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	db := db.Open("sqlite3", "./dev.sqlite3.db")

	a := app.New(db)
	a.RegisterHandlers(r)

	fmt.Println("listening...")
	panic(http.ListenAndServe(":8080", r))
}
