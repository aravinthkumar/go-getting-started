package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// Method
// To differentiate a method from function is by providing a variable before the method name
// This method also implements the http/Handler interface since it matches the inteface signature
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller"))
}

// Constructor Function
// Note : convention should be new<Object>
func newUserController() *userController {
	// With Structs we can send the addressOf operator &userController
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
