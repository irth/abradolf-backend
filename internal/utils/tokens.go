package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
func GenerateRandomToken(bytes int) (string, error) {
	b := make([]byte, bytes)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}
