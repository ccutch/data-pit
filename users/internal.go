package users

import (
	"context"

	"github.com/ccutch/data-pit/models"
	"github.com/ccutch/datastore-model"
)

// GenerateTestUser creates a test user and returns a response of that user
func GenerateTestUser(ctx context.Context) interface{} {
	u := &models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	db.NewDatastore(ctx).Create(u)
	return NewUserResponse(ctx, u)
}
