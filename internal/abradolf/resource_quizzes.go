package abradolf

import (
	"net/http"

	"github.com/gorilla/mux"
)

type QuizzesResource struct{}

func (q *QuizzesResource) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("", q.List).Methods("GET")
	r.HandleFunc("", q.Create).Methods("POST")

	r.HandleFunc("/{id}", q.Get).Methods("GET", "DELETE")

	r.HandleFunc("/{id}/questions", q.ListQuestions).Methods("GET")
	r.HandleFunc("/{id}/questions", q.CreateQuestion).Methods("POST")

	r.HandleFunc("/{id}/questions/{qid}", q.GetQuestion).Methods("GET", "DELETE")
}

func (q *QuizzesResource) List(w http.ResponseWriter, r *http.Request) {
	// TODO: implement QuizzesResource.List
	w.Write([]byte("list"))
}

func (q *QuizzesResource) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: implement QuizzesResource.Create
	w.Write([]byte("create"))
}

func (q *QuizzesResource) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: implement QuizzesResource.Get
	w.Write([]byte("get"))
}

func (q *QuizzesResource) ListQuestions(w http.ResponseWriter, r *http.Request) {
	// TODO: implement QuizzesResource.ListQuestions
	w.Write([]byte("listQuestions"))
}

func (q *QuizzesResource) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: implement QuizzesResource.CreateQuestion
	w.Write([]byte("createQuestion"))
}

func (q *QuizzesResource) GetQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: implement QuizzesResource.GetQuestion
	w.Write([]byte("getQuestion"))
}
