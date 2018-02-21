package main

import (
	"resto_rest_api/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {

	router := mux.NewRouter()
	router.HandleFunc("/restaurants", GetRestaurants).Methods("GET")

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

func GetReservations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantId, _ := strconv.Atoi(vars["id"])

	var reservations = repository.GetReservations(restaurantId)
	json.NewEncoder(w).Encode(reservations)
}

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantId, _ := strconv.Atoi(vars["id"])

	type B struct {
		ReservationTime string `json:"reservation_time"`
	}

	var b B

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	layout := "2006-01-02 15:04:05"
	reservationTime, _ := time.Parse(layout, b.ReservationTime)

	repository.CreateReservation(restaurantId, reservationTime);

	respondWithJSON(w, http.StatusCreated, b);
}


func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}