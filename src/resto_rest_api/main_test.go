package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
)

var a App

func TestGetRestaurants(t *testing.T) {

	a = App{}
	a.Initialize()

	req, _ := http.NewRequest("GET", "/restaurants", nil)

	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetReservations(t *testing.T) {

	a = App{}
	a.Initialize()

	req, _ := http.NewRequest("GET", "/restaurants/3098/reservations", nil)

	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateReservation(t *testing.T) {

	a = App{}
	a.Initialize()

	payload := []byte(`{"reservation_time":"2018-02-21 02:00:00"}`);

	req, _ := http.NewRequest("POST", "/restaurants/3098/reservations", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
