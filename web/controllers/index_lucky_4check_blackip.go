package controllers

import (
	"Lottery-go/models"
	"Lottery-go/services"
	"time"
)

func (api *LuckyApi) checkBlackIp(ip string) (bool, *models.Blackip) {
	info := services.NewBlackIpService().GetByIp(ip)
	if info == nil || info.Ip == "" {
		return false, nil
	}
	if info.BlackTime > int(time.Now().Unix()) {
		return true, info
	}
	return false, nil
}
