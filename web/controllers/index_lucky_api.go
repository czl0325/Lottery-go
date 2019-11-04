package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/conf"
	"Lottery-go/models"
	"Lottery-go/services"
	"Lottery-go/web/utils"
	"log"
	"time"
)

type LuckyApi struct {
}

func (api *LuckyApi) luckyDo(uid int, username, ip string) (int, string, *models.ObjGiftPrize) {
	// 2 用户抽奖分布式锁定
	ok := utils.LockLucky(uid)
	if ok {
		defer utils.UnlockLucky(uid)
	} else {
		return 101, "正在抽奖，请稍后", nil
	}

	// 3 验证用户今日参与次数
	userDayNum := utils.IncrUserLuckyNum(uid)
	if userDayNum > conf.UserPrizeMax {
		return 102, "今日的抽奖次数已用完，明天再来吧", nil
	} else {
		ok = api.checkUserDay(uid, userDayNum)
		if !ok {
			return 103, "今日的抽奖次数已用完，明天再来吧", nil
		}
	}

	// 4 验证IP今日的参与次数
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		return 104, "相同IP参与次数太多，明天再来参与吧", nil
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

	// 8 匹配奖品是否中奖
	prizeGift := api.prize(prizeCode, limitBlack)
	if prizeGift == nil || prizeGift.PrizeNum <= 0 || (prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		return 201, "很遗憾，没有中奖，请下次再试", nil
	}

	// 9 有限制奖品发放
	if prizeGift.Id > 0 {
		if utils.GetGiftPoolNum(prizeGift.Id) <= 0 {
			return 202, "很遗憾，没有中奖，请下次再试", nil
		}
		ok := utils.PrizeGift(prizeGift.Id)
		if ok == false {
			return 203, "很遗憾，没有中奖，请下次再试", nil
		}
	}

	//10 不同编码的优惠券的发放

	//11 记录中奖记录
	result := &models.Result{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        uid,
		UserName:   username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: int(time.Now().Unix()),
		SysIp:      ip,
		SysStatus:  0,
	}
	err := services.NewResultService().Create(result)
	if err != nil {
		log.Println("保存中奖纪录失败,", err)
		return 209, "很遗憾，没有中奖，请下次再试", nil
	}
	return 0, "成功抽到奖品", prizeGift
}
