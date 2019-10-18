package controllers

import (
	"Lottery-go/models"
	"Lottery-go/services"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminResultController struct {
	Ctx 			iris.Context
	ServiceUser    	services.UserService
	ServiceGift    	services.GiftService
	ServiceCode    	services.CodeService
	ServiceResult  	services.ResultService
	ServiceUserDay 	services.UserdayService
	ServiceBlackIp 	services.BlackIpService
}

func (c *AdminResultController) Get() mvc.Result {
	giftId := c.Ctx.URLParamIntDefault("gift_id",0)
	uid := c.Ctx.URLParamIntDefault("uid", 0)
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""
	var dataList []models.Result
	if giftId > 0 {
		dataList = c.ServiceResult.SearchByGift(giftId, page, size)
	} else if uid > 0 {
		dataList = c.ServiceResult.SearchByUser(uid, page, size)
	} else {
		dataList = c.ServiceResult.GetAll(page, size)
	}
	var total int64
	if giftId > 0 {
		total = c.ServiceResult.CountByGift(giftId)
	} else if uid > 0 {
		total = c.ServiceResult.CountByUser(uid)
	} else {
		total = c.ServiceResult.CountAll()
	}
	if len(dataList) >= size {
		pageNext = fmt.Sprintf("%d",page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d",page-1)
	}
	return mvc.View{
		Name: "admin/result.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "result",
			"GiftId":   giftId,
			"Uid":      uid,
			"DataList": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminResultController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceResult.Delete(id)
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		 refer = "/admin/result"
	}
	return mvc.Response{
		Path: refer,
	}
}

func (c *AdminResultController) GetCheat() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceResult.Update(&models.Result{Id:id, SysStatus:0},[]string{"sys_status"})
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/result"
	}
	return mvc.Response{
		Path: refer,
	}
}