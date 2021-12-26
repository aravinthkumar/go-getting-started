package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {

	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new creation should not include id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {

	for _, v := range users {
		if id == v.ID {
			return *v, nil
		}
	}
	return User{}, errors.New("User doesn't exist")
}

func UpdateUser(u User) (User, error) {
	for i, v := range users {
		if u.ID == v.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, errors.New("User not found to update")

}

func DeleteUser(id int) error {
	for i, v := range users {
		if id == v.ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User '%v'  deleted", id)
}
