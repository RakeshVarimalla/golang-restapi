package vesp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	vesp = "vesp"
)
func Vesp( c *gin.Context){
	c.String(http.StatusOK,vesp)
}