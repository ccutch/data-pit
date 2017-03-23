package api

import (
	"context"
	"net/http"
)

// Responder gives a response for the api
type Responder interface {
	Respond(context.Context, http.ResponseWriter)
}

// ContextResponder takes a context and returns a responder
type ContextResponder func(context.Context) interface{}
