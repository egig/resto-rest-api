package model

type Restaurant struct {
	ID  string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude  float32 `json:"longitude,omitempty"`
}