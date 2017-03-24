package users

/**
 * UserResponse follows the api.Responder interface of the api package
 * using response classes to format, sanitize, add virtual fields, and
 * normalize data from models.
 */

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/ccutch/data-pit/models"
	"google.golang.org/appengine/user"
)

// UserResponse is a generalized model which should be returned by all functions
// in the users package and be servable from the api.
type UserResponse struct {
	// User model, gets all fields from this model
	*models.User
	err   error
	Token string `json:"token"`
}

// NewUserResponse creates a new user response instance.
// Handles hiding private fields when user making request is not user
// being requested and adds private fields when the user being requested
// is the user making the request.
func NewUserResponse(ctx context.Context, u *models.User) *UserResponse {
	user := user.Current(ctx)
	res := &UserResponse{User: u}
	isUser := user != nil && user.Email == u.Email

	if isUser {
		// Add private data
		res.Token = "A valid token"
	} else {
		// Remove private data
		res.Email = ""
	}
	return res
}

// JSON formats user response as a string of JSON
func (u *UserResponse) JSON() string {
	b, _ := json.Marshal(u)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	return string(out.String())
}

// String overrides print behavior for easier logging
func (u *UserResponse) String() string {
	return u.JSON()
}

// Error returns the error if one is present
func (u *UserResponse) Error() error {
	return u.err
}
