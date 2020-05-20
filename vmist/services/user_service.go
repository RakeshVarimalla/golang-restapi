package services

import (
	"github.com/RakeshVarimalla/golang-restapi/vmist/domain"
	"github.com/RakeshVarimalla/golang-restapi/vmist/utils"
)

func GetUser(userId int64) (*domain.User, *utils.VmistError) {

	return domain.GetUser(userId)
}
