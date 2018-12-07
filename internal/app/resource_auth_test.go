package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/irth/abradolf-backend/internal/app"

	"github.com/irth/abradolf-backend/internal/db"
	"github.com/irth/abradolf-backend/internal/db/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setup() (a *app.AuthResource, d *gorm.DB) {
	d = db.Open("sqlite3", ":memory:")
	a = &app.AuthResource{DB: d}
	return
}

func postRequest(handler http.HandlerFunc, path string, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, path, strings.NewReader(body))
	handler(w, r)
	return w
}

type errResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func TestRegisterAccount(t *testing.T) {
	u := map[string]string{
		"username": "testuser",
		"password": "password",
	}

	body, _ := json.Marshal(u)

	a, db := setup()
	defer db.Close()

	w := postRequest(a.Register, "/register", string(body))

	require.Equal(t, http.StatusCreated, w.Code, "the response code should equal 201")

	var users []models.User
	db.Where("username = ?", u["username"]).Find(&users)

	require.Equal(t, 1, len(users), "only one user should be created")
	require.True(t, users[0].CheckPassword(u["password"]), "the password in the created record should match the requested password")
}

func TestRegisterAccountWithEmptyUsername(t *testing.T) {
	a, db := setup()
	defer db.Close()

	body := `{"username":"","password":"123"}`
	w := postRequest(a.Register, "/register", body)
	require.Equal(t, http.StatusUnprocessableEntity, w.Code, "the response code should equal 422")

	var e errResponse
	json.NewDecoder(w.Body).Decode(&e)
	require.Equal(t, app.ErrUsernameEmpty, e.Code, "it should return a correct error code")

	var c int
	db.Model(&models.User{}).Count(&c)
	require.Equal(t, 0, c, "no user should be created")
}

func TestRegisterAccountTwice(t *testing.T) {
	body := `{
		"username": "testuser",
		"password": "testpassword"
	}`

	a, db := setup()
	defer db.Close()

	w1 := postRequest(a.Register, "/register", body)
	require.Equal(t, http.StatusCreated, w1.Code, "the response code should equal 201")

	w2 := postRequest(a.Register, "/register", body)
	require.Equal(t, http.StatusUnprocessableEntity, w2.Code, "the response code should equal 422")

	var e errResponse
	json.NewDecoder(w2.Body).Decode(&e)
	require.Equal(t, app.ErrUsernameTaken, e.Code, "it should return a correct error code")

	var c int
	db.Model(&models.User{}).Count(&c)
	require.Equal(t, 1, c, "only one user should be created")
}
