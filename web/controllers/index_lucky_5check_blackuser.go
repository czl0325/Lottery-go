package controllers

import (
	"Lottery-go/models"
	"Lottery-go/services"
	"time"
)

func (api *LuckyApi) checkBlackUser(uid int) (bool, *models.User) {
	info := services.NewUserService().Get(uid)
	if info != nil && info.BlackTime > int(time.Now().Unix()) {
		return true, info
	}
	return false, info
}
