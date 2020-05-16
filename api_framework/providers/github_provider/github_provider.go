package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"microservices/api/clients/restclient"
	"microservices/api/domain/github"
	"net/http"
)

const (
	headerAutherization      = "Authorization"
	headerAutherizationToken = "token %s"
	urlCreateRepo            = "https://api.github.com/user/repos"
)

func GetAutherzationHeader(accessToken string) string {
	return fmt.Sprintf(headerAutherizationToken, accessToken)
}
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	//Autherization: token 4a3b7e42ccb60dfa5ce7e881b64aa95f9e62cf3f
	headers := http.Header{}
	log.Println(GetAutherzationHeader(accessToken))
	log.Println(urlCreateRepo)
	headers.Set(headerAutherization, GetAutherzationHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying create repo in github: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		err := json.Unmarshal(bytes, &errResponse)
		if err != nil {
			return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	err1 := json.Unmarshal(bytes, &result)
	if err1 != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response %s", err1.Error()))
		return nil, &github.GitHubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal create repo successful response"}

	}
	return &result, nil
}
