package oauth

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservices/api/utils/errors"
	"microservices/oauth-api/src/api/domain/oauth"
	"microservices/oauth-api/src/api/services"
	"net/http"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}
	log.Println(request)
	token, err1 := services.OauthService.CreateAccessToken(&request)
	if err1 != nil {
		c.JSON(err1.Status(), err1)
	}
	c.JSON(http.StatusCreated,token)
}
func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")
	token, err := services.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.Status(), err)
	}
	c.JSON(http.StatusOK,token)
}
