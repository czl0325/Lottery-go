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
		ok = api.checkUserDay(uid, userDayNum)
		if !ok {
			return int(comm.NoLotteryNum), "今日的抽奖次数已用完，明天再来吧", nil
		}
	}

	// 4 验证IP今日的参与次数
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		return int(comm.IpLimit), "相同IP参与次数太多，明天再来参与吧", nil
	}

	limitBlack := false
	if ipDayNum > conf.IpPrizeMax {
		limitBlack = true
	}

	// 5 验证IP黑名单
	var blackInfo *models.Blackip
	if !limitBlack {
		limitBlack, blackInfo = api.checkBlackIp(ip)
	}

	// 6 验证用户黑名单
	var blackUser *models.User
	if !limitBlack {
		limitBlack, blackUser = api.checkBlackUser(uid)
	}

	// 7 获得抽奖编码
	prizeCode := comm.Random(10000)

}
