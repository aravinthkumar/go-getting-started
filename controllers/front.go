package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()
	// http.Handle function expects pattern and a handler since uc implements the ServeHTTP interface it returns the http.Handler
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
