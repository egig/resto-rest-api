package main_test

import (
	"."
	"net/http"
	"net/http/httptest"
	"testing"
)

var a main.App

func TestGetRestaurants(t *testing.T) {

	a = main.App{}
	a.Initialize()

	req, _ := http.NewRequest("GET", "/restaurants", nil)

	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
