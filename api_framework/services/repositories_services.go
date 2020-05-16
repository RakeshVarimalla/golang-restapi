package services

import (
	"log"
	"microservices/api/config"
	"microservices/api/domain/github"
	"microservices/api/domain/repositories"
	"microservices/api/providers/github_provider"
	"microservices/api/utils/errors"
	"strings"
)

type reposService struct{

}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse,errors.ApiErrors)
}
var (
	RepositoryService reposServiceInterface
)

func init(){
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse,errors.ApiErrors){
input.Name = strings.TrimSpace(input.Name)
log.Println("services.CreateRepo")
if input.Name == ""{
	return nil,errors.NewBadRequestError("invalid repository name")
}

request := github.CreateRepoRequest{
	Name: input.Name,
	Description: input.Description,
	Private :false,
}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(),request)
	if err != nil{
		return nil, errors.NewApiError(err.StatusCode,err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id :  response.Id,
		Name: response.Name,
		Owner: response.Owner.Login,
	}
return &result,nil
}