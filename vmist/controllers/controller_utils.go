package controllers

import "C"
import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, statusCode int, body interface{}){
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(statusCode,body)
		return
	}
	c.JSON(statusCode, body)
}
