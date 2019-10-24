package controllers

import (
	"Lottery-go/services"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminCodeController struct {
	Ctx         iris.Context
	ServiceCode services.CodeService
}

func (c *AdminCodeController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	size := 100
	pagePrev := ""
	pageNext := ""
	dataList := c.ServiceCode.GetAll(page, size)
	total := len(dataList)
	if total >= size {
		total = int(c.ServiceCode.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}
	return mvc.View{
		Name: "/admin/code.html",
		Data: iris.Map{
			"Title":    "优惠券管理",
			"Channel":  "code",
			"GiftId":   giftId,
			"DataList": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}
