package app

import (
	"context"
	"fmt"
	"github.com/irth/abradolf-backend/internal/db/models"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

func NewAuthMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := strings.TrimSpace(r.Header.Get("Authorization"))
			split := strings.Split(tokenHeader, " ")
			if len(split) != 2 || strings.ToLower(split[0]) != "bearer" {
				next.ServeHTTP(w, r)
				return
			}

			token := split[1]
			var authToken models.AuthToken
			notFound := db.Find(&authToken, models.AuthToken{Token: token}).RecordNotFound()
			if notFound {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "user", authToken.UserID)
			fmt.Println(authToken.UserID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
