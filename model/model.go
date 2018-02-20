package model

import "time"

type Restaurant struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type Reservation struct {
	ID              string      `json:"id,omitempty"`
	ReservationTime time.Time   `json:"reservation_time,omitempty"`
	Restaurant      *Restaurant `json:"restaurant,omitempty"`
}
