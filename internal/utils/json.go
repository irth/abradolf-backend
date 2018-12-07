package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

const ErrInvalidJSON ErrCode = "ERR_INVALID_JSON"

func UnmarshalBody(w http.ResponseWriter, body io.ReadCloser, target interface{}) (err error) {
	err = json.NewDecoder(body).Decode(target)
	if err != nil {
		WriteErrorString(w, http.StatusBadRequest, ErrInvalidJSON, "An error occured while parsing the request.")
	}
	return
}

func WriteJSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(obj)
}

type H = map[string]interface{}
