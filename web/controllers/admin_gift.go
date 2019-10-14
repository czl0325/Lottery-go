package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/services"
	"Lottery-go/web/utils"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminGiftController struct {
	Ctx 			iris.Context
	ServiceUser    	services.UserService
	ServiceGift    	services.GiftService
	ServiceCode    	services.CodeService
	ServiceResult  	services.ResultService
	ServiceUserday 	services.UserdayService
	ServiceBlackip 	services.BlackIpService
}

func (c *AdminGiftController) Get() mvc.Result {
	dataList := c.ServiceGift.GetAll(false)
	for i, gift := range dataList {
		// 奖品发放的计划数据
		prizeData := make([][2]int, 0)
		err := json.Unmarshal([]byte(gift.PrizeData), &prizeData)
		if err != nil || prizeData == nil || len(prizeData) < 1 {
			dataList[i].PrizeData = "[]"
		} else {
			newpd := make([]string, len(prizeData))
			for index, pd := range prizeData {
				ct := comm.FormatFromUnixTime(int64(pd[0]))
				newpd[index] = fmt.Sprintf("【%s】: %d", ct , pd[1])
			}
			str, err := json.Marshal(newpd)
			if err == nil && len(str) > 0 {
				dataList[i].PrizeData = string(str)
			} else {
				dataList[i].PrizeData = "[]"
			}
		}
		// 奖品当前的奖品池数量
		num := utils.GetGiftPoolNum(gift.Id)
		dataList[i].Title = fmt.Sprintf("【%d】%s", num, dataList[i].Title)
	}
	total := len(dataList)
	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title" : "奖品管理后台",
			"Channel" : "gift",
			"DataList" : dataList,
			"Total" : total,
		},
		Layout: "admin/layout.html",
	}
}