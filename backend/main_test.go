package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router().ServeHTTP(w, req)

	expected := "OK"
	value := w.Body.String()
	if expected != value {
		t.Fatalf("Expected %s but got %s", expected, value)
	}
}
