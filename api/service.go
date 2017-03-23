package api

import (
	"context"
	"net/http"
)

// Service is an http handler tied to a url which follows REST conventions
// and returns a predicted response
type Service interface {
	Name() string
	Methods() map[string]ContextResponder
	Response() Responder
}

// DefaultService provides a template and fallback for other services
type DefaultService struct{}

// Name fulfils service interface
func (ds *DefaultService) Name() string {
	return "Default Service"
}

// Methods fulfils service interface
func (ds *DefaultService) Methods() map[string]ContextResponder {
	return map[string]ContextResponder{
		"GET":    methodNotImplemented,
		"POST":   methodNotImplemented,
		"PUT":    methodNotImplemented,
		"DELETE": methodNotImplemented,
	}
}

// Helper Context responder for methods that arent implemented
type notImplementedResponder struct{}

func (r *notImplementedResponder) Respond(_ context.Context, w http.ResponseWriter) {
	http.Error(w, "Method not implemented", http.StatusNotImplemented)
}

func methodNotImplemented(ctx context.Context) interface{} {
	return new(notImplementedResponder)
}
