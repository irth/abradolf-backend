package models

import (
	"time"

	"github.com/irth/abradolf-backend/internal/utils"

	"github.com/jinzhu/gorm"
)

type AuthToken struct {
	gorm.Model

	User    User      `json:"-"`
	UserID  uint      `json:"-"`
	Token   string    `json:"auth_token"`
	Expires time.Time `json:"expires"`
}

func NewAuthToken(u User) (*AuthToken, error) {
	t, err := utils.GenerateRandomToken(64)

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
