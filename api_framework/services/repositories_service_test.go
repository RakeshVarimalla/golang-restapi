package services

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"microservices/api/clients/restclient"
	"microservices/api/domain/repositories"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMocks()
	os.Exit(m.Run())
}
 func TestReposService_CreateRepoInValidInputName(t *testing.T) {
	 request := repositories.CreateRepoRequest{}
	 res, err := RepositoryService.CreateRepo(request)
	 assert.Nil(t, res)
	 assert.NotNil(t, err)
	 assert.EqualValues(t, "invalid repository name", err.Message())
	 assert.EqualValues(t,http.StatusBadRequest,err.Status())
 }

 func TestReposService_CreateRepoErrFromGithub(t *testing.T) {
 	restclient.FlushMocks()
 	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})
	 request := repositories.CreateRepoRequest{
	 	Name:"testing",
	 }
	 res, err := RepositoryService.CreateRepo(request)
 	assert.Nil(t,res)
	 assert.NotNil(t,err)
	 assert.EqualValues(t,http.StatusUnauthorized,err.Status())
	 assert.EqualValues(t,"Requires authentication",err.Message())

 }

func TestReposService_CreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{ "id":123,"name":"testing","owner":{"login":"rakesh"}}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name:"testing"}
	res, err := RepositoryService.CreateRepo(request)
	assert.Nil(t,err)
	assert.NotNil(t,res)
	assert.EqualValues(t,123,res.Id)
	assert.EqualValues(t,"testing",res.Name)
	assert.EqualValues(t,"rakesh",res.Owner)

}