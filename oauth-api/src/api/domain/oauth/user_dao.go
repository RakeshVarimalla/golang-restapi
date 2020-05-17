package oauth

import (
	"log"
	"microservices/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? and password=?"
)

var (
	users = map[string]*User{
		"rakesh": &User{
			Id:       123,
			Username: "rakesh",
		},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiErrors) {
	log.Println(username)
	log.Println(password)
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}
