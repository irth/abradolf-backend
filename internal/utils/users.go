package utils

import (
	"net/http"

	"github.com/irth/abradolf-backend/internal/db/models"
	"github.com/jinzhu/gorm"
)

func GetUser(db *gorm.DB, r *http.Request) *models.User {
	uid, ok := r.Context().Value("user").(uint)
	if !ok {
		return nil
	}

	var user models.User
	if db.First(&user, uid).Error != nil {
		return nil
	}

	return &user
}
