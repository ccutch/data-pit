package models

import "github.com/ccutch/datastore-model"

// User model to store information about the user
type User struct {
	db.Model
	Email string `json:"email"`
	Name  string `json:"name"`
}
