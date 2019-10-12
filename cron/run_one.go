package cron

import (
	"Lottery-go/comm"
	"Lottery-go/services"
	"log"
)

func ConfigueAppOneCron()  {

}

// 重置所有奖品的发奖计划
// 每5分钟执行一次
func resetAllGiftPrizeData() {
	giftService := services.NewGiftService()
	list := giftService.GetAll(false)
	nowTime := comm.NowUnix()
	for _, gift := range list {
		if gift.PrizeTime != 0 && (gift.PrizeData == "" || gift.PrizeEnd <= nowTime) {
			// 立即执行
			log.Println("crontab start utils.ResetGiftPrizeData giftInfo=", gift)

		}
	}
}