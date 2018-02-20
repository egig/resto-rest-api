package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/restaurants", GetRestaurants).Methods("GET")
  router.HandleFunc("/restaurants/{id}", GetRestaurant).Methods("GET")
  router.HandleFunc("/restaurants/{id}", CreateRestaurant).Methods("POST")
  router.HandleFunc("/restaurants/{id}", DeleteRestaurant).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
  var restaurants []Restaurant

  restaurants = append(restaurants, Restaurant{ID: "1", Name: "Common Ground", Latitude: -6.211110, Longitude: 106.816363});
  restaurants = append(restaurants, Restaurant{ID: "2", Name: "Common Ground 2", Latitude: -6.749354, Longitude: 106.837483});
  restaurants = append(restaurants, Restaurant{ID: "3", Name: "Common Ground 3", Latitude: -6.545410, Longitude: 106.349433});
  json.NewEncoder(w).Encode(restaurants);
}
func GetRestaurant(w http.ResponseWriter, r *http.Request) {}
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {}
func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {}

type Restaurant struct {
  ID  string `json:"id,omitempty"`
  Name  string `json:"name,omitempty"`
  Latitude  float32 `json:"latitude,omitempty"`
  Longitude  float32 `json:"longitude,omitempty"`
}
