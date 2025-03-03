package main

import (
	"audio_converter/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHealthCheckHandler проверяет, что эндпоинт `/health` работает
func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	expected := "Server is running!"
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("Expected response %q, got %q", expected, actual)
	}
}
