package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFound_ServeHTTP(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/notfound", nil)
	resp := httptest.NewRecorder()

	h := NotFound{}
	h.ServeHTTP(resp, req)

	got := resp.Result().StatusCode
	want := http.StatusNotFound
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
