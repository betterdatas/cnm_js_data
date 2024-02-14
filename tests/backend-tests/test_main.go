// tests/backend-tests/test_main.go

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Home handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Welcome to Code Notation Manager"
	if rr.Body.String() != expected {
		t.Errorf("Home handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCNMAPIHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cnm", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cnmAPIHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CNM API handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Code Notation Manager API"
	if rr.Body.String() != expected {
		t.Errorf("CNM API handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func main() {
	fmt.Println("Running backend tests...")
	t.Run("TestHomeHandler", TestHomeHandler)
	t.Run("TestCNMAPIHandler", TestCNMAPIHandler)
}
