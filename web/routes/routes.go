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

	index := mvc.New(b.Party("/"))
	index.Register(userService,
		giftService,
		codeService,
		resultService,
		userDayService,
		blackIpService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Register(userService,
		giftService,
		codeService,
		resultService,
		userDayService,
		blackIpService)
	admin.Handle(new(controllers.AdminController))
}
