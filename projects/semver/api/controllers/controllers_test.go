package controllers

import (
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/josepht96/learning/projects/semver/api/middleware"
)

func testDefaultGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
  
	if err != nil {
	  t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.Handler(newDefaultHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func testProjectGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/project/", nil)
  
	if err != nil {
	  t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.Handler(newProjectHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func testProjectPut(t *testing.T) {
	req, err := http.NewRequest("PUT", "/project/", nil)
  
	if err != nil {
	  t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.Handler(newProjectHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func testProjectPost(t *testing.T) {
	req, err := http.NewRequest("POST", "/project/", nil)
  
	if err != nil {
	  t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.Handler(newProjectHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func testProjectsGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/projects", nil)
  
	if err != nil {
	  t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.Handler(newProjectsHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}


func testUserGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/", nil)
  
	if err != nil {
	  t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.Handler(newUserHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}