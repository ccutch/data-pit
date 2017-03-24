package users

/**
 * Because we want a large set of data and because we want to
 * be able to test user methods, a random user generator is
 * going to be helpful.
 */

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/ccutch/data-pit/models"
	"github.com/icrowley/fake"
	"googlemaps.github.io/maps"
)

const (
	mapsAPIKey = "AIzaSyAF8Yi7P6MCw9-iLmqD7mSbXvJUuSLWGEc"
)

var (
	mapsClient *maps.Client
	mockCities = []string{
		"San Fransisco, California",
		"New York City, New York",
		"Chicago, Illinois",
		"Miami, Flordia",
		"Austin, Texas",
	}
)

func init() {
	mapsClient, _ = maps.NewClient(maps.WithAPIKey(mapsAPIKey))
}

// createMockLocation creates a mock location for mock users
func createMockLocation(ctx context.Context, client *datastore.Client) *datastore.Key {
	// Getting a mock city from predefined list
	city := mockCities[rand.Intn(len(mockCities))]

	// Geocode lat and long from google maps information
	req := &maps.GeocodingRequest{
		Address: city,
	}
	res, err := mapsClient.Geocode(ctx, req)
	if err != nil {
		fmt.Println("Error getting city", city, res, err)
		return nil
	}

	// Create instance of Location model
	loc := &models.Location{
		Latitude:   res[0].Geometry.Location.Lat,
		Longitude:  res[0].Geometry.Location.Lng,
		PlaceID:    res[0].PlaceID,
		StreetName: res[0].FormattedAddress,
	}

	// Write location to database
	locKey := datastore.IncompleteKey("Location", nil)
	locKey, err = client.Put(ctx, locKey, loc)

	if err != nil {
		return nil
	}
	return locKey
}

// CreateTestUser creats a test user with mock data.
// This is much faster than thinking up a user.
func CreateTestUser() *UserResponse {
	// Open connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "data-pit")
	if err != nil {
		return &UserResponse{err: err}
	}

	defer client.Close()

	// Create instance of user model
	name := fake.FirstName() + " " + fake.LastName()
	email := strings.Replace(strings.ToLower(name), " ", "-", -1) + "@test.com"
	user := &models.User{
		Created:     time.Now(),
		Updated:     time.Now(),
		Email:       email,
		Name:        name,
		LocationKey: createMockLocation(ctx, client),
	}

	// Write user to datastore
	key := datastore.IncompleteKey("User", nil)
	_, err = client.Put(ctx, key, user)
	if err != nil {
		return &UserResponse{err: err}
	}

	return &UserResponse{User: user}
}
