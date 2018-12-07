package utils

import (
	"net/http"
)

type ErrCode string

func WriteErrorString(w http.ResponseWriter, httpCode int, errCode ErrCode, err string) {
	WriteJSON(w, httpCode, map[string]string{
		"error": err,
		"code":  string(errCode),
	})
}

func WriteError(w http.ResponseWriter, httpCode int, errCode ErrCode, err error) {
	WriteErrorString(w, httpCode, errCode, err.Error())
}
