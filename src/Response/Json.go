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

func ToStruct(w http.ResponseWriter, r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	JsonError(w, err)
}

func JsonError(w http.ResponseWriter, err error) {
	if err != nil {
		payload := ErrorResponse{
			Error:   true,
			Message: err.Error(),
		}
		ToJson(w, http.StatusBadRequest, payload)
	}
}
