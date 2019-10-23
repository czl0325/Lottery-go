package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/services"
	"Lottery-go/web/viewmodels"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminUserController struct {
	Ctx 			iris.Context
	ServiceUser    	services.UserService
	ServiceGift    	services.GiftService
	ServiceCode    	services.CodeService
	ServiceResult  	services.ResultService
	ServiceUserDay 	services.UserdayService
	ServiceBlackIp 	services.BlackIpService
}

func (c* AdminUserController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""

	dataList := c.ServiceUser.GetAll(page, size)
	total := c.ServiceUser.CountAll()
	if len(dataList) < int(total) {
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d",page-1)
	}
	return mvc.View{
		Name: "admin/user.html",
		Data: iris.Map{
			"Title":"管理后台",
			"Channel":  "user",
			"DataList": dataList,
			"Total":    total,
			"Now":      comm.NowUnix(),
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout:"admin/layout.html",
	}
}

func (c* AdminUserController) GetEdit() mvc.Result {
	return mvc.View{
		Name:"admin/userEdit.html",
		Data:iris.Map{
			"Title":"添加用户",
			"Channel":"user",
		},
		Layout:"admin/layout.html",
	}
}

func (c* AdminUserController) PostSave() mvc.Result {
	data := viewmodels.ViewUser{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		fmt.Println("读取用户信息填写有误", err)
		return mvc.Response{
			Text: fmt.Sprintf("ReadForm转换异常, err=%s", err),
		}
	}
	return mvc.View{

	}
}