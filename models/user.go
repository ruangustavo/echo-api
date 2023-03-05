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
	users      = make(map[int]*User)
	nextUserID = 0
)

func GetAllUsers() []User {
	res := make([]User, 0, len(users))

	for _, user := range users {
		res = append(res, *user)
	}

	return res
}

func GetUserById(id int) (*User, error) {
	user, ok := users[id]

	if !ok {
		return nil, errors.NewUserError("User not found")
	}

	return user, nil
}

func CreateUser(user *User) error {
	user.ID = nextUserID
	users[user.ID] = user
	nextUserID++
	return nil
}

func UpdateUser(user *User) error {
	_, ok := users[user.ID]

	if !ok {
		return errors.NewUserError("User not found")
	}

	users[user.ID] = user
	return nil
}

func DeleteUser(user *User) error {
	_, ok := users[user.ID]

	if !ok {
		return errors.NewUserError("User not found")
	}

	delete(users, user.ID)
	return nil
}
