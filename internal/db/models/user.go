package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model

	Username     string
	PasswordHash []byte
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = hash
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil
}
