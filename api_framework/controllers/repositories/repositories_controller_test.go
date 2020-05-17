package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"microservices/api/clients/restclient"
	"microservices/api/domain/repositories"
	"microservices/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMocks()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request,_ := http.NewRequest(http.MethodConnect,"/repositories",strings.NewReader(``))
	c.Request = request
	CreateRepo(c)
	assert.EqualValues(t,http.StatusBadRequest,response.Code)
	fmt.Println(response.Body.String())
	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.NotNil(t,apiErr)
	assert.Nil(t,err)
	assert.EqualValues(t,http.StatusBadRequest,apiErr.Status())
	assert.EqualValues(t,"invalid json body",apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request,_ := http.NewRequest(http.MethodConnect,"/repositories",strings.NewReader(`{"name":"testing"}`))
	c.Request = request

	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	CreateRepo(c)
	assert.EqualValues(t,http.StatusUnauthorized,response.Code)
	fmt.Println(response.Body.String())
	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.NotNil(t,apiErr)
	assert.Nil(t,err)
	assert.EqualValues(t,http.StatusUnauthorized,apiErr.Status())
	assert.EqualValues(t,"Requires authentication",apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request,_ := http.NewRequest(http.MethodConnect,"/repositories",strings.NewReader(`{"name":"testing"}`))
	c.Request = request

	restclient.FlushMocks()
	restclient.AddMocks(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	CreateRepo(c)
	assert.EqualValues(t,http.StatusCreated,response.Code)
	fmt.Println(response.Body.String())
	var res repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(),&res)
	assert.Nil(t,err)
	assert.EqualValues(t,123,res.Id)
}

