package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	code int8
	error
	message string
	method  bool
}

func (e *Error) response(w http.ResponseWriter) {
	if e.error != nil {
		data := map[string]interface{}{"code": e.code, "message": e.message, "error": e.error}
		if e.method {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		} else {
			fmt.Print(data)
		}
		return
	}
}
