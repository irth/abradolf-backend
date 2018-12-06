package abradolf

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Abradolf struct{}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func (a Abradolf) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/", greet)
}
