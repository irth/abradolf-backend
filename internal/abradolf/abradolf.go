package abradolf

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Abradolf struct {
	Auth AuthResource

	Quizzes QuizzesResource
	Users   UsersResource
}

func New() *Abradolf {
	return &Abradolf{
		Auth: AuthResource{},

		Quizzes: QuizzesResource{},
		Users:   UsersResource{},
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func (a Abradolf) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/", greet)

	a.Auth.RegisterHandlers(r)

	a.Quizzes.RegisterHandlers(r.PathPrefix("/quizzes").Subrouter())
	a.Users.RegisterHandlers(r.PathPrefix("/users").Subrouter())
}
