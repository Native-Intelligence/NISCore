package core

import (
	"encoding/json"
	"net/http"
)

type Service struct {
}

func (s *Service) response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func internalResponse() {

}
