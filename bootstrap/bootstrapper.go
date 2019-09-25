package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

type Configurator func(bootstrapper *Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions *sessions.Sessions
}

func NewApp(appName, appOwner string, cfgs ...Configurator) *Bootstrapper  {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Sessions:     nil,
	}

	for _,cfg := range cfgs{
		cfg(b)
	}
	return b
}