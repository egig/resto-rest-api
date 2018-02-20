package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"./repository"
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

	a.Router = router
}

func (a *App) Run(port string) {
	log.Println("Running application on port", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	var restaurants = repository.GetRestaurants();
	json.NewEncoder(w).Encode(restaurants);
}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {}
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {}
func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {}