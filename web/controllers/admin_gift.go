package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/models"
	"Lottery-go/services"
	"Lottery-go/web/utils"
	"Lottery-go/web/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"time"
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

func (c *AdminGiftController) GetEdit() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	gift := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(id, false)
		if data != nil {
			gift.Id = data.Id
			gift.Title = data.Title
			gift.PrizeNum = data.PrizeNum
			gift.PrizeCode = data.PrizeCode
			gift.PrizeTime = data.PrizeTime
			gift.Img = data.Img
			gift.DisplayOrder = data.DisplayOrder
			gift.Gtype = data.Gtype
			gift.Gdata = data.Gdata
			gift.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
			gift.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
		}
	}
	return mvc.View{
		Name: "admin/giftEdit.html",
		Data: iris.Map{
			"Title" : "奖品编辑",
			"Channel" : "gift",
			"info" : gift,
		},
		Layout: "admin/layout.html",
	}
}

func (c* AdminGiftController) GetResult() mvc.Result {
	return mvc.View{
		Name: "admin/result.html",
		Data: iris.Map{
			"Title" : "奖品编辑",
			"Channel" : "gift",
		},
		Layout: "admin/layout.html",
	}
}

func (c* AdminGiftController) PostSave() mvc.Result {
	data := viewmodels.ViewGift{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		fmt.Println("读取保存奖品数据失败", err)
		return mvc.Response{
			Text: fmt.Sprintf("ReadForm转换异常, err=%s", err),
		}
	}
	gift := models.Gift{}
	gift.Id = data.Id
	gift.Title = data.Title
	gift.PrizeNum = data.PrizeNum
	gift.PrizeCode = data.PrizeCode
	gift.PrizeTime = data.PrizeTime
	gift.Img = data.Img
	gift.DisplayOrder = data.DisplayOrder
	gift.Gtype = data.Gtype
	gift.Gdata = data.Gdata
	t1, err1 := comm.ParseTime(data.TimeBegin)
	t2, err2 := comm.ParseTime(data.TimeEnd)
	if err1 != nil || err2 != nil {
		return mvc.Response{
			Text: fmt.Sprintf("开始时间、结束时间的格式不正确, err1=%s, err2=%s", err1, err2),
		}
	}
	gift.TimeBegin = int(t1.Unix())
	gift.TimeEnd = int(t2.Unix())
	if gift.Id > 0 {
		data2 := c.ServiceGift.Get(gift.Id, false)
		if data2 != nil {
			gift.SysUpdated = int(time.Now().Unix())
			gift.SysIp = comm.ClientIP(c.Ctx.Request())
			// 对比修改的内容项
			if gift.PrizeNum != data2.PrizeNum {
				gift.LeftNum = data2.LeftNum - data2.PrizeNum - gift.PrizeNum
				if gift.LeftNum < 0 || gift.PrizeNum <= 0 {
					gift.LeftNum = 0
				}
				gift.SysStatus = data2.SysStatus
				utils.ResetGiftPrizeData(&gift, c.ServiceGift)
			} else {
				gift.LeftNum = gift.PrizeNum
			}
			if gift.PrizeTime != data2.PrizeTime {
				// 发奖周期发生了变化
				utils.ResetGiftPrizeData(&gift, c.ServiceGift)
			}
			c.ServiceGift.Update(&gift, []string{"title", "prize_num", "left_num", "prize_code", "prize_time", "img", "display_order", "gtype", "gdata", "time_begin", "time_end", "sys_updated"})
		} else {
			gift.Id = 0
		}
	}
	if gift.Id == 0 {
		gift.LeftNum = gift.PrizeNum
		gift.SysIp = comm.ClientIP(c.Ctx.Request())
		gift.SysCreated = int(time.Now().Unix())
		gift.SysUpdated = int(time.Now().Unix())
		c.ServiceGift.Create(&gift)
		// 更新奖品的发奖计划
		utils.ResetGiftPrizeData(&gift, c.ServiceGift)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}

func (c* AdminGiftController) GetDelete() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	if id > 0 {
		c.ServiceGift.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}