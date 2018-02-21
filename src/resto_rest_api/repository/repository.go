package repository

import (
	"resto_rest_api/config"
	"resto_rest_api/model"
)
import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func CreateReservation(restaurantId int, reservationTime time.Time) sql.Result {
	c := config.GetValue();
	db, _ := sql.Open(c.DBDriver, c.DBDSN);
	query := `INSERT INTO reservations(reservation_time, restaurant_id) VALUES(?,?)`

	res, _ := db.Exec(query, reservationTime.Format("2006-01-02 15:04:05") , restaurantId)
	return res
}


func GetReservations(restaurantId int) []model.Reservation {

	c := config.GetValue();
	db, _ := sql.Open(c.DBDriver, c.DBDSN);

	var restaurant model.Restaurant
	query1 := `SELECT id,name,latitude,longitude,district_name,cuisine_name,rating FROM restaurants WHERE id=?`

	err := db.QueryRow(query1, restaurantId).Scan(
		&restaurant.ID,
		&restaurant.Name,
		&restaurant.Latitude,
		&restaurant.Longitude,
		&restaurant.DistrictName,
		&restaurant.CuisineName,
		&restaurant.Rating)

	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		log.Fatal(err)
	default:
	}

	query := `SELECT id,reservation_time FROM reservations WHERE restaurant_id=?`
	rows, err := db.Query(query, restaurantId)

	if err != nil {
		log.Fatal(err)
	}

	var reservations []model.Reservation

	defer rows.Close()
	for rows.Next() {
		var r model.Reservation
		err := rows.Scan(
			&r.ID,
			&r.ReservationTime)

		if err != nil {
			log.Fatal(err)
		}

		r.Restaurant  = &restaurant

		reservations = append(reservations, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


	return reservations
}


func GetRestaurants(latitude, longitude float64) []model.Restaurant {

	c := config.GetValue();
	db, err := sql.Open(c.DBDriver, c.DBDSN);
	query := `SELECT id,name,latitude,longitude,district_name,cuisine_name,rating,
				(
					6371 *
					acos(
						cos( radians( ? ) ) *
						cos( radians( latitude ) ) *
						cos(
							radians(longitude ) - radians( ? )
						) +
						sin(radians(?)) *
						sin(radians(latitude))
					)
				) as distance
				FROM restaurants
				HAVING
					distance < ?
				ORDER BY
					distance ASC`

	rows, err := db.Query(query, latitude, longitude, latitude, 5)
	if err != nil {
		log.Fatal(err)
	}

	var restaurants []model.Restaurant

	defer rows.Close()
	for rows.Next() {
		var r model.Restaurant
		err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Latitude,
			&r.Longitude,
			&r.DistrictName,
			&r.CuisineName,
			&r.Rating,
			&r.Distance)
		if err != nil {
			log.Fatal(err)
		}

		restaurants = append(restaurants, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


	return restaurants
}
