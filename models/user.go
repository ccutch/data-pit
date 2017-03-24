package models

import (
	"time"

	"cloud.google.com/go/datastore"
)

// User model to store information about the user
type User struct {
	Created     time.Time      `json:"created"`
	Updated     time.Time      `json:"updated"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	LocationKey *datastore.Key `json:"-"`
}
