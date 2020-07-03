package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/baburay23/src/handlers"
	"github.com/go-playground/assert/v2"
)

//func TestXxx(*testing.T)

func TestEducationRoute(t *testing.T) {
	//r := gin.Default()

	r := setupServer()
	w := httptest.NewRecorder()
	//NewRecorder returns an initialized ResponseRecorder.
	req, _ := http.NewRequest("GET", "/education", nil)

	//NewRequest returns a new incoming server Request, suitable for passing to an http.Handler for testing.
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Bored API":"Learn Rust"}`, w.Body.String())

}

func TestCharityRoute(t *testing.T) {
	//r := gin.Default()

	r := setupServer()
	w := httptest.NewRecorder()
	//NewRecorder returns an initialized ResponseRecorder.
	req, _ := http.NewRequest("GET", "/charity", nil)
	//NewRequest returns a new incoming server Request, suitable for passing to an http.Handler for testing.
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Bored API":"Volunteer at foodbank"}`, w.Body.String())

}
