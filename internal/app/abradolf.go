package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/irth/abradolf-backend/internal/utils"
	"github.com/jinzhu/gorm"
)

type Abradolf struct {
	Auth AuthResource

	Quizzes QuizzesResource
	Users   UsersResource

	db *gorm.DB
}

func New(db *gorm.DB) *Abradolf {
	return &Abradolf{
		Auth: AuthResource{DB: db},

		Quizzes: QuizzesResource{},
		Users:   UsersResource{},

		db: db,
	}
}

func (a Abradolf) Greet(w http.ResponseWriter, r *http.Request) {
	u := utils.GetUser(a.db, r)
	un := "nil"
	if u != nil {
		un = u.Username
	}

	fmt.Fprintf(w, "Hello World! %s, u=%s", time.Now(), un)
}

func (a Abradolf) RegisterHandlers(r *mux.Router) {
	r.Use(NewAuthMiddleware(a.db))
	r.HandleFunc("/", a.Greet)

	a.Auth.RegisterHandlers(r)

	a.Quizzes.RegisterHandlers(r.PathPrefix("/quizzes").Subrouter())
	a.Users.RegisterHandlers(r.PathPrefix("/users").Subrouter())
}
