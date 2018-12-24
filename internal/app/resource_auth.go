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

type authRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthResource) Login(w http.ResponseWriter, r *http.Request) {
	var body authRequestBody

	if utils.UnmarshalBody(w, r.Body, &body) != nil {
		return
	}

	var u models.User
	result := a.DB.Where("username = ?", body.Username).First(&u)

	if result.RecordNotFound() || !u.CheckPassword(body.Password) {
		utils.WriteErrorString(w, http.StatusUnauthorized, ErrUnauthorized, "Incorrect username or password.")
		return
	}

	if result.Error != nil {
		utils.WriteErrorString(w, http.StatusInternalServerError, ErrDatabaseError, "Database error.")
		return
	}

	t, err := models.NewAuthToken(u)
	if err != nil {
		utils.WriteErrorString(w, http.StatusInternalServerError, ErrUnknown, "Internal server error.")
		return
	}

	err = a.DB.Create(t).Error
	if err != nil {
		utils.WriteErrorString(w, http.StatusInternalServerError, ErrDatabaseError, "Database error.")
		return
	}

	utils.WriteJSON(w, http.StatusOK, t)
}

func (a *AuthResource) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: implement AuthResource.Logout
	w.Write([]byte("logout"))
}

func (a *AuthResource) Register(w http.ResponseWriter, r *http.Request) {
	var body authRequestBody

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
