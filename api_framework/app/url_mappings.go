package app

import (
	"microservices/api/controllers/repositories"
	"microservices/api/controllers/vesp"
)

func mapUrls(){
	router.GET("/vesp",vesp.Vesp)
	router.POST("/repositories",repositories.CreateRepo)

}
