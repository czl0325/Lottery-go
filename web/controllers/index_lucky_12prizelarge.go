package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/models"
	"Lottery-go/services"
)

func (api *LuckyApi) prizeLarge(ip string, uid int, username string, userInfo *models.User, blackInfo *models.Blackip) {
	userService := services.NewUserService()
	blackService := services.NewBlackIpService()
	nowTime := comm.NowUnix()
	blackTime := 30 * 86400
	// 更新用户的黑名单信息
	if userInfo == nil || userInfo.Id <= 0 {
		userInfo = &models.User{
			Id:         uid,
			UserName:   username,
			BlackTime:  nowTime + blackTime,
			SysCreated: nowTime,
			SysUpdated: nowTime,
			SysIp:      ip,
		}
		userService.Create(userInfo)
	} else {
		userInfo = &models.User{
			Id:         uid,
			UserName:   username,
			BlackTime:  nowTime + blackTime,
			SysUpdated: nowTime,
			SysIp:      ip,
		}
		userService.Update(userInfo, nil)
	}

	if blackInfo != nil || blackInfo.Id <= 0 {
		blackInfo = &models.Blackip{
			Ip:         ip,
			BlackTime:  nowTime + blackTime,
			SysCreated: nowTime,
			SysUpdated: nowTime,
		}
		blackService.Create(blackInfo)
	} else {
		blackInfo.BlackTime = nowTime + blackTime
		blackInfo.SysUpdated = nowTime
		blackService.Update(blackInfo, nil)
	}
}
