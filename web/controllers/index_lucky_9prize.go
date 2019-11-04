package controllers

import (
	"Lottery-go/conf"
	"Lottery-go/models"
	"Lottery-go/services"
)

func (api *LuckyApi) prize(prizeCode int, limitBlack bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	giftList := services.NewGiftService().GetAllUse(true)
	for _, gift := range giftList {
		if gift.PrizeCodeA <= prizeCode &&
			gift.PrizeCodeB >= prizeCode {
			if !limitBlack || gift.Gtype < conf.GtypeGiftSmall {
				prizeGift = &gift
			}
		}
	}
	return prizeGift
}
