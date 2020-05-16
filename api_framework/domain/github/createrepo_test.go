package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T){
	apptest := CreateRepoRequest{
			Name: "Hello-World-Code",
			Description: "This is your first repository from code",
			Homepage: "https://github.com",
			Private: false,
			HasIssues: true,
			HasProjects: true,
			HasWiki: true,
	}
	bytes, err := json.Marshal(apptest)
	assert.Nil(t,err)
	assert.NotNil(t,bytes)
	fmt.Println(string(bytes))
	var testjson CreateRepoRequest
	json.Unmarshal(bytes,&testjson)
	assert.NotNil(t,testjson)
	assert.EqualValues(t,apptest.Name, testjson.Name)
	fmt.Println(testjson)

}