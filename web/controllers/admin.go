package controllers

import (
	"Lottery-go/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx 		   iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserDay services.UserdayService
	ServiceBlackIp services.BlackIpService
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name:"admin/index.html",
		Data:iris.Map{
			"Title": "陈昭良",
			"Channel":"gift",
		},
		Layout:"admin/layout.html",
	}
}