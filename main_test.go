package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()
	
	// Check if response contains HTML structure
	if !strings.Contains(body, "<html>") {
		t.Error("Response should contain HTML structure")
	}
	
	// Count occurrences of "Hello World"
	count := strings.Count(body, "Hello World")
	if count != 25 {
		t.Errorf("Expected 25 occurrences of 'Hello World', got %d", count)
	}
}