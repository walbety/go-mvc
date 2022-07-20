package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {
			Id: 1, FirstName: "Frodo", LastName: "Baggins",
		},
	}
)

func GetUser(userId int64) (*User, error) {

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("cant find user: %v", userId))

}
