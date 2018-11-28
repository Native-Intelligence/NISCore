package core

import (
	"gopkg.in/validator.v2"
	"net/http"
)

type Request struct {
}

var ex Error

func (r *Request) validate(w http.ResponseWriter, class interface{}) {
	ex.error = validator.Validate(class)
	ex.response(w)
	// println(result)
}
