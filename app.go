package main

import (
	"./repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {

	router := mux.NewRouter()
	router.HandleFunc("/restaurants", GetRestaurants).Methods("GET")
	router.HandleFunc("/restaurants/{id}", GetRestaurant).Methods("GET")
	router.HandleFunc("/restaurants/{id}", CreateRestaurant).Methods("POST")
	router.HandleFunc("/restaurants/{id}", DeleteRestaurant).Methods("DELETE")

	router.HandleFunc("/restaurants/{id}/reservations", GetReservations).Methods("GET")
	router.HandleFunc("/restaurants/{id}/reservations", CreateReservation).Methods("POST")

	a.Router = router
}

func (a *App) Run(port string) {
	log.Println("Running application on port", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	latitude, _ := strconv.ParseFloat(r.FormValue("latitude"), 32)
	longitude, _ := strconv.ParseFloat(r.FormValue("longitude"), 32)
	var restaurants = repository.GetRestaurants(latitude, longitude)
	json.NewEncoder(w).Encode(restaurants)
}

func GetRestaurant(w http.ResponseWriter, r *http.Request)    {}
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {}
func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {}

func GetReservations(w http.ResponseWriter, r *http.Request)   {}
func CreateReservation(w http.ResponseWriter, r *http.Request) {}
