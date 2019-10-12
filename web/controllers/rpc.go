package controllers

import (
	"Lottery-go/services"
	"github.com/kataras/iris"
)

type RpcController struct {
	Ctx 		   *iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserDay services.UserdayService
	ServiceBlackIp services.BlackIpService
}

