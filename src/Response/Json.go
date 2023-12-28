package Response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func ToJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	JsonError(w, err)
}
