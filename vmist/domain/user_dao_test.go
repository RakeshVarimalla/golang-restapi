package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
    //Initialization

	//Execution
	user, err := GetUser(0)
	/*
		if user != nil{
			t.Error("We were not expecting user with id 0")
		}
	*/

	//Validation
	assert.Nil(t, user, "we were not expecting a user with id 0")
	/*
		if err == nil{
			t.Error("We were expecting error when user id 0")
		}
		if err.StatusCode != http.StatusNotFound{
		t.Error("We were expecting 404 when user is not found")
		}
	*/
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}
func TestGetUserNoError(t *testing.T) {
	var userdata User
	userdata.Id = 123;
	userdata.FirstName = "Rakesh"
	userdata.LastName = "Varimalla"
	userdata.Email = "Rakeshv@vitalpointz.net"
	AddUser(userdata)

	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Rakesh", user.FirstName)
	assert.EqualValues(t, "Varimalla", user.LastName)
	assert.EqualValues(t, "Rakeshv@vitalpointz.net", user.Email)

}

func BenchmarkAddUser(b *testing.B) {
	var userdata User
	userdata.Id = 1;
	userdata.FirstName = "Rakesh"
	userdata.LastName = "Varimalla"
	userdata.Email = "rakeshv@vitalpointz.net"
	for i:=0; i<b.N; i++ {
		userdata.Id = uint64(i)
		AddUser(userdata)
	}

}
func BenchmarkGetUser(b *testing.B) {
for i:=0; i< b.N; i++ {
	GetUser(int64(b.N))
}
}

