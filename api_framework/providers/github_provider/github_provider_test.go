package github_provider

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"microservices/api/clients/restclient"
	"microservices/api/domain/github"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMocks()
	os.Exit(m.Run())

}
func TestConstants(t *testing.T){
	assert.EqualValues(t,"Authorization",headerAutherization)
	assert.EqualValues(t,"token %s",headerAutherizationToken)
	assert.EqualValues(t,"https://api.github.com/user/repos",urlCreateRepo)
}
func TestGetAutherzationHeader(t *testing.T) {
	s := GetAutherzationHeader("4a3b7e42ccb60dfa5ce7e881b64aa95f9e62cf3f")
	assert.EqualValues(t, "token 4a3b7e42ccb60dfa5ce7e881b64aa95f9e62cf3f", s)
}
func TestCreateRepoErrorRestclient(t *testing.T) {
    restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err : errors.New("invalid rest client response"),
	})
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,resp)
	assert.NotNil(t,err)
	assert.EqualValues(t,"invalid rest client response", err.Message)


}
func TestCreateRepoInvalidResponseBody(t *testing.T) {
	inValidCloser, _ := os.Open("_-asf3")
    restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: inValidCloser,
		},
	})
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,resp)
	assert.NotNil(t,err)
	assert.EqualValues(t,"invalid response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
		},
	})
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,resp)
	assert.NotNil(t,err)
	assert.EqualValues(t,"invalid json body", err.Message)
}


func TestCreateRepoUnAutherized(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,resp)
	assert.NotNil(t,err)
	assert.EqualValues(t,"Requires authentication", err.Message)
}


func TestCreateInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": "263952452","node_id": "MDEwOlJlcG9zaXRvcnkyNjM5NTI0NTI=","name": "Hello-World-Test-auth","full_name": "RakeshVarimalla/Hello-World-Test-auth"}
`)),
		},
	})
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,resp)
	assert.NotNil(t,err)
	//assert.EqualValues(t,263952452, resp.ID)
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 263952452,"node_id": "MDEwOlJlcG9zaXRvcnkyNjM5NTI0NTI=","name": "Hello-World-Test-auth","full_name": "RakeshVarimalla/Hello-World-Test-auth"}`)),
		},
	})
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,err)
	assert.NotNil(t,resp)
	assert.EqualValues(t,263952452, resp.Id)
}