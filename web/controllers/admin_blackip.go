package controllers

import (
	"Lottery-go/services"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminBlackIpController struct {
	Ctx            iris.Context
	ServiceBlackIp services.BlackIpService
}

func (c *AdminBlackIpController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	dataList := c.ServiceBlackIp.GetAll(page, size)
	total := len(dataList)
	if total >= size {
		total = int(c.ServiceBlackIp.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	return mvc.View{
		Name: "/admin/blackip.html",
		Data: iris.Map{
			"Title":    "黑名单管理",
			"Channel":  "blackip",
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}
