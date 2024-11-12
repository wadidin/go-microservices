package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)

	// Create a response recorder to record the response from the handler.
	rr := httptest.NewRecorder()

	// Call the handler with the request and recorder.
	handler(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Hello, World v1b"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
