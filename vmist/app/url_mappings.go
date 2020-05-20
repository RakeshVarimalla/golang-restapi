package app

import (
	"github.com/RakeshVarimalla/golang-restapi/vmist/controllers"
)

func mapUrls(){
	router.GET("/users/:user_id", controllers.GetUser)
}
