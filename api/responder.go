package api

import "context"

// Responder gives a response for the api
type Responder interface{}

// ContextResponder takes a context and returns a responder
// TODO: why wont this work?
// 			 type ContextResponder func(context.Context) Responder
type ContextResponder func(context.Context) interface{}
