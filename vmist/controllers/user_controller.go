package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/RakeshVarimalla/golang-restapi/vmist/services"
	"github.com/RakeshVarimalla/golang-restapi/vmist/utils"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		userErr := &utils.VmistError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		//c.JSON(http.StatusBadRequest,userErr)
		Respond(c,http.StatusBadRequest,userErr)
		return
	}
	user, userErr := services.GetUser(userId)
	if userErr != nil {
		//c.JSON(userErr.StatusCode,userErr)
		Respond(c,userErr.StatusCode,userErr)
		return
	}
	// return user to client
	//c.JSON(http.StatusOK,user)
	Respond(c,http.StatusOK,user)
}
