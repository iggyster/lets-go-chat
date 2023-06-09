package handler

import (
	"net/http"
)

type NotFound struct{}

func (handler *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
