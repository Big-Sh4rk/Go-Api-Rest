package controllers

import (
	"net/http"
	"testing"
)

func TestGetBook(t *testing.T) {
	response, err := http.Get("http://localhost:8080/book/")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK,
			response.StatusCode)
	}
	if err != nil {
		t.Errorf("Encountered an error:", err)
	}
}

func TestGetBookById(t *testing.T) {
	response, err := http.Get("http://localhost:8080/book/1")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK,
			response.StatusCode)
	}
	if err != nil {
		t.Errorf("Encountered an error:", err)
	}
}

func TestCreateBook(t *testing.T) {
	var json = []byte(`{"Name":"Zero to Hero", "Author":"Phil Collin", "Publication":"Amazon"}`)
	response, err := http.Post("http://localhost:8080/book/")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK,
			response.StatusCode)
	}
	if err != nil {
		t.Errorf("Encountered an error:", err)
	}

}

func TestUpdateBook(t *testing.T) {

}

func TestDeleteBook(t *testing.T) {
	response, err := http.Get("http://localhost:8080/book/1")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK,
			response.StatusCode)
	}
	if err != nil {
		t.Errorf("Encountered an error:", err)
	}
}
