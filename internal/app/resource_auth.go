package app

import (
	"net/http"

	"github.com/irth/abradolf-backend/internal/db/models"
	"github.com/irth/abradolf-backend/internal/utils"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type AuthResource struct {
	DB *gorm.DB
}

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
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if utils.UnmarshalBody(w, r.Body, &body) != nil {
		return
	}

	if len(body.Username) == 0 {
		utils.WriteErrorString(w, http.StatusUnprocessableEntity, ErrUsernameEmpty, "username cannot be empty")
		return
	}

	u := models.User{
		Username: body.Username,
	}

	u.SetPassword(body.Password)

	tx := a.DB.Begin()
	var c int
	tx.Model(&models.User{}).Where("username = ?", u.Username).Count(&c)
	if c != 0 {
		tx.Rollback()
		utils.WriteErrorString(w, http.StatusUnprocessableEntity, ErrUsernameTaken, "username already taken")
		return
	}

	err := tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		utils.WriteErrorString(w, http.StatusInternalServerError, ErrDatabaseError, "database error")
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		utils.WriteErrorString(w, http.StatusInternalServerError, ErrDatabaseError, "database error")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
