package handler

import (
	"encoding/json"
	"net/http"
)

type Errors struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (errors *Errors) AddError(field, message, detail string) {
	errors.Errors = append(errors.Errors, &Error{Field: field, Message: message, Detail: detail})
}

func (errors *Errors) Count() int {
	return len(errors.Errors)
}

func SendErrors(resp http.ResponseWriter, errors Errors, status int) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(errors)
	if err != nil {
		http.Error(resp, "Failed encoding", http.StatusInternalServerError)
	}
}
