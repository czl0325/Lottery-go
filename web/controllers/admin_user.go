package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/models"
	"Lottery-go/services"
	"Lottery-go/web/viewmodels"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"time"
)

type AdminUserController struct {
	Ctx 			iris.Context
	ServiceUser    	services.UserService
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
	t, err := comm.ParseTime(data.BlackTime)
	if err != nil {
		return mvc.Response{
			Text: fmt.Sprintf("黑名单到期时间格式有误, err=%s", err),
		}
	}
	user := models.User{
		UserName:   data.UserName,
		BlackTime:  int(t.Unix()),
		RealName:   data.RealName,
		Mobile:     data.Mobile,
		Address:    data.Address,
		SysCreated: int(time.Now().Unix()),
		SysUpdated: int(time.Now().Unix()),
		SysIp:      comm.ClientIP(c.Ctx.Request()),
	}
	err = c.ServiceUser.Create(&user)
	if err != nil {
		return mvc.Response{
			Text: fmt.Sprintf("添加用户失败, err=%s", err),
		}
	}
	return mvc.Response {
		Path: "/admin/user",
	}
}

func (c* AdminUserController) GetBlack() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	t := c.Ctx.URLParamIntDefault("time",0)
	if id > 0 {
		user := c.ServiceUser.Get(id)
		user.BlackTime = int(time.Now().Unix()) + t * 86400
		err := c.ServiceUser.Update(user, []string{"black_time", "sys_update"})
		if err != nil {
			return mvc.Response{
				Text: fmt.Sprintf("设置用户黑名单失败, err=%s", err),
			}
		}
	}
	return mvc.Response{
		Path: "/admin/user",
	}
}