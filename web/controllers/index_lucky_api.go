package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/conf"
	"Lottery-go/models"
	"Lottery-go/web/utils"
)

type LuckyApi struct {
}

func (api *LuckyApi) luckyDo(uid int, username, ip string) (int, string, *models.ObjGiftPrize) {
	// 2 用户抽奖分布式锁定
	ok := utils.LockLucky(uid)
	if ok {
		defer utils.UnlockLucky(uid)
	} else {
		return int(comm.NowLottery), "正在抽奖，请稍后", nil
	}

	// 3 验证用户今日参与次数
	userDayNum := utils.IncrUserLuckyNum(uid)
	if userDayNum > conf.UserPrizeMax {
		return int(comm.NoLotteryNum), "今日的抽奖次数已用完，明天再来吧", nil
	} else {

	}
}
