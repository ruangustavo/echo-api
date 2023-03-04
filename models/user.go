package models

import (
	"echo-api/controllers/errors"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var (
	userList []User
	id       int
)

func GetAllUsers() []User {
	if len(userList) == 0 {
		return []User{}
	}

	return userList
}

func GetUserById(id int) (*User, error) {
	for _, user := range userList {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.NewUserError("User not found")
}

func CreateUser(user *User) error {
	user.ID = id
	userList = append(userList, *user)
	id++

	return nil
}

func UpdateUser(user *User) error {
	for i, u := range userList {
		if user.ID == u.ID {
			userList[i] = *user
			return nil
		}
	}

	return errors.NewUserError("User not found")
}

func DeleteUser(user *User) error {
	for i, u := range userList {
		if user.ID == u.ID {
			userList = append(userList[:i], userList[i+1:]...)
			return nil
		}
	}

	return errors.NewUserError("User not found")
}
