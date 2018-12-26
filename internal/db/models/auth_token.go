package models

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/jinzhu/gorm"
)

type AuthToken struct {
	gorm.Model

	User    User      `json:"-"`
	UserID  uint      `json:"-"`
	Token   string    `json:"auth_token"`
	Expires time.Time `json:"expires"`
}

// https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
func generateRandomToken(bytes int) (string, error) {
	b := make([]byte, bytes)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

func NewAuthToken(u User) (*AuthToken, error) {
	t, err := generateRandomToken(64)

	if err != nil {
		return nil, err
	}

	return &AuthToken{
		User:    u,
		UserID:  u.ID,
		Token:   t,
		Expires: time.Now().Add(time.Hour * 24 * 30 * 6),
	}, nil
}
