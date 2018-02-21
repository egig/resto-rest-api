package repository

import (
	"resto_rest_api/config"
	"resto_rest_api/model"
)
import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)


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
