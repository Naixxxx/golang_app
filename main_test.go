package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code is 200, got %d", rec.Code)
	}

	expected := "Hello World"
	body := rec.Body.String()

	if strings.TrimSpace(body) != strings.TrimSpace(expected) {
		t.Errorf("expected body is %q, got %q", expected, body)
	}
}
