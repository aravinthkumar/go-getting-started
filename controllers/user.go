package controllers

import (
	"encoding/json"
	"github/aravinthkumar/go-getting-started/models"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// Method
// To differentiate a method from function is by providing a variable before the method name
// This method also implements the http/Handler interface since it matches the inteface signature
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader((http.StatusNotFound))
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(w, id)
		case http.MethodDelete:
			uc.delete(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		}
	}

}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJson(models.GetUsers(), w)

}

func (uc *userController) get(w http.ResponseWriter, id int) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJson(u, w)

}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := parseResquestToUser(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := parseResquestToUser(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}

func (uc *userController) delete(id int, w http.ResponseWriter) {

	err := models.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}

func parseResquestToUser(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil

}

func encodeResponseAsJson(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

// Constructor Function
// Note : convention should be new<Object>
func newUserController() *userController {
	// With Structs we can send the addressOf operator &userController
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
