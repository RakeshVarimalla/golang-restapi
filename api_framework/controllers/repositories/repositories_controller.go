package repositories

import (
	"github.com/gin-gonic/gin"
	"microservices/api/domain/repositories"
	"microservices/api/services"
	"microservices/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context){

	var request repositories.CreateRepoRequest

	err := c.ShouldBindJSON(&request)
	if err != nil{
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(),apiErr)
		return
	}
	result, err1 := services.RepositoryService.CreateRepo(request)
	if err1 != nil {
		c.JSON(err1.Status(),err1)
		return
	}
	c.JSON(http.StatusCreated,result)
}

