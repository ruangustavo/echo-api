package models

import "fmt"

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
		userList = append(userList, User{ID: id, Name: "John", Email: "johndoe@email.com"})
		id = 1
	}

	return userList
}

func GetUserById(id int) (*User, error) {
	for _, user := range userList {
		if user.ID == id {
			return &user, nil
		}
	}

	err := fmt.Errorf("User not found")
	return nil, err
}

func CreateUser(user *User) error {
	user.ID = id
	userList = append(userList, *user)
	id++

	return nil
}
