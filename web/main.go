package main

import (
	"Lottery-go/bootstrap"
	"Lottery-go/web/middleware/identity"
	"Lottery-go/web/routes"
)

func newApp() *bootstrap.Bootstrapper  {
	app := bootstrap.NewApp("抽奖系统","陈昭良")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main()  {
	app := newApp()
	app.Listen("localhost:8080")
}
