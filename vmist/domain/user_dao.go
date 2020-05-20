package domain

import (
	"fmt"
	"net/http"
	"github.com/RakeshVarimalla/golang-restapi/vmist/utils"
)

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
//var users  map[int64]*User
/*
var users = make(map[int64]*User)
func init() {
	user2 := User{
		Id: 12345,
		FirstName: "Shashikanth",
		LastName: "Varimalla",
		Email: "Shashikanth@gmail.com",
	}
	users[int64(user2.Id)] = &user2
	users[1234] = &User{Id: 1234, FirstName: "Rakesh", LastName: "Varimalla", Email: "Rakeshv@vitalpointz.net"}
}
*/
var users = map[int64]*User{
	123:  {Id: 123, FirstName: "Rakesh", LastName: "Varimalla", Email: "Rakeshv@vitalpointz.net"},
	1234: {Id: 1234, FirstName: "Rakesh", LastName: "Varimalla", Email: "Rakeshv@vitalpointz.net"},
}

func AddUser(userdata User){
	users[int64(userdata.Id)] = &userdata
}

func GetUser(userId int64) (*User, *utils.VmistError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.VmistError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
