package services

import (
	"vmist/domain"
	"vmist/utils"
)

func GetUser(userId int64) (*domain.User, *utils.VmistError) {

	return domain.GetUser(userId)
}
