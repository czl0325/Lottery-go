package routes

import (
	"Lottery-go/bootstrap"
	"Lottery-go/services"
	"Lottery-go/web/controllers"
	"github.com/kataras/iris/mvc"
)

func Configure(b* bootstrap.Bootstrapper)  {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userDayService := services.NewUserdayService()
	blackIpService := services.NewBlackIpService()

	rpc := mvc.New(b.Party("/"))
	rpc.Register(userService,
		giftService,
		codeService,
		resultService,
		userDayService,
		blackIpService)
	rpc.Handle(new(controllers.IndexController))
}
