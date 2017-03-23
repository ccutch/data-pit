/**
 * UserResponse follows the api.Responder interface of the api package
 * using response classes to format, sanitize, add virtual fields, and
 * normalize data from models.
 */
package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ccutch/data-pit/models"
	"google.golang.org/appengine/user"
)

// UserResponse is a generalized model which should be returned by all functions
// in the users package and be servable from the api.
type UserResponse struct {
	// User model, gets all fields from this model
	*models.User

	// Virual fields that are generated for the response but do not exist in the
	// database
	Token string `json:"token"`
}

// NewUserResponse creates a new user response instance.
// Handles hiding private fields when user making request is not user
// being requested and adds private fields when the user being requested
// is the user making the request.
func NewUserResponse(ctx context.Context, u *models.User) UserResponse {
	user := user.Current(ctx)
	res := UserResponse{User: u}
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

// Respond writes user response to http response writer
func (u UserResponse) Respond(ctx context.Context, w http.ResponseWriter) {
	e := json.NewEncoder(w)
	e.Encode(u)
}
