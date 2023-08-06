package utils

import (
	"net/http"
)

func Respond(w http.ResponseWriter, status int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
