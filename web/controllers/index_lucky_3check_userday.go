package controllers

import (
	"Lottery-go/conf"
	"Lottery-go/models"
	"Lottery-go/services"
	"Lottery-go/web/utils"
	"fmt"
	"log"
	"time"
)

func (api *LuckyApi) checkUserDay(uid int, num int64) bool {
	userDayService := services.NewUserdayService()
	userDayInfo := userDayService.GetUserToday(uid)
	if userDayInfo != nil && userDayInfo.Uid == uid {
		// 今天存在抽奖记录
		if userDayInfo.Num >= conf.UserPrizeMax {
			if int(num) < userDayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userDayInfo.Num))
			}
			return false
		} else {
			userDayInfo.Num++
			if int(num) < userDayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userDayInfo.Num))
			}
			err := userDayService.Update(userDayInfo, nil)
			if err != nil {
				log.Println("更新用户每日抽奖次数失败=", err103)
			}
		}
	} else {
		y, m, d := time.Now().Date()
		strDay := fmt.Sprintf("%04d%02d%02d", y, m, d)
		userDayInfo = &models.Userday{
			Uid:        uid,
			Day:        strDay,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
			SysUpdated: int(time.Now().Unix()),
		}
		err := userDayService.Create(userDayInfo)
		if err != nil {
			if err != nil {
				log.Println("创建用户每日抽奖次数失败=", err103)
			}
		}
		utils.InitUserLuckyNum(uid, 1)
	}
	return true
}
