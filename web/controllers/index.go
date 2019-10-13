package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/models"
	"Lottery-go/services"
	"fmt"
	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx 		   iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserDay services.UserdayService
	ServiceBlackIp services.BlackIpService
}

func (c* IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "欢迎进入抽奖系统,<a href='/public/index.html'>开始抽奖</a>"
}

func (c *IndexController) GetGifts() map[string]interface{}  {
	rs := make(map[string]interface{})
	dataList := c.ServiceGift.GetAll(false)
	list := make([]models.Gift, 0)
	for _, gift := range dataList {
		// 正常状态的才需要放进来
		if gift.SysStatus == 0 {
			list = append(list, gift)
		}
	}
	rs["code"] = 0
	rs["msg"] = ""
	rs["data"] = list
	return rs
}

func (c* IndexController) GetLogin() {
	uid := comm.Random(10000)
	loginUser := &models.ObjLoginUser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), loginUser)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index?from=loginIn")
}

func (c* IndexController) GetLoginOut() {
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index?from=loginOut")
}