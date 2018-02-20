package repository

import "../model"

func GetRestaurants() []model.Restaurant {

	var restaurants []model.Restaurant

	restaurants = append(restaurants, model.Restaurant{ID: "1", Name: "Common Ground", Latitude: -6.211110, Longitude: 106.816363})
	restaurants = append(restaurants, model.Restaurant{ID: "2", Name: "Common Ground 2", Latitude: -6.749354, Longitude: 106.837483})
	restaurants = append(restaurants, model.Restaurant{ID: "3", Name: "Common Ground 3", Latitude: -6.545410, Longitude: 106.349433})
	return restaurants
}
