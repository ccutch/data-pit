package models

// Location model stores data from google maps api
type Location struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	StreetName string  `json:"street_name"`
	PlaceID    string  `json:"place_id"`
}
