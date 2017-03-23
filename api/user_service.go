package api

import "github.com/ccutch/data-pit/users"

// UserService /api/user
type UserService struct {
	DefaultService
}

// Name fulfils service interface
func (us *UserService) Name() string {
	return "User Service"
}

// Response fulfils service interface
func (us *UserService) Response() Responder {
	return new(users.UserResponse)
}

// Methods fulfils service interface
func (us *UserService) Methods() map[string]ContextResponder {
	return map[string]ContextResponder{
		"GET":  users.GetListOfUsers,
		"POST": users.CreateUser,
	}
}

// AdditionalRoutes adds more routes
func (us *UserService) AdditionalRoutes() map[string]ContextResponder {
	return map[string]ContextResponder{
		"/generate-user": users.GenerateTestUser,
		"/signin":        users.Signin,
		"/signup":        users.Signup,
	}
}
