package app

import (
	"microservices/api/controllers/vesp"
	"microservices/oauth-api/src/api/controllers/oauth"
)

func mapUrls(){
	router.GET("/vesp",vesp.Vesp)
	router.POST("/oauth/access_token",oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id",oauth.GetAccessToken)

}


