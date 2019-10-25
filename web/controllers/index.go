package controllers

import (
	"Lottery-go/comm"
	"Lottery-go/conf"
	"Lottery-go/models"
	"Lottery-go/services"
	"fmt"
	"github.com/kataras/iris"
	"strconv"
	"time"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserDay services.UserdayService
	ServiceBlackIp services.BlackIpService
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "欢迎进入抽奖系统,<a href='/public/index.html'>开始抽奖</a>"
}

func (c *IndexController) GetGifts() map[string]interface{} {
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

func (c *IndexController) GetLogin() {
	uid := comm.Random(10000)
	loginUser := &models.ObjLoginUser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginUser(c.Ctx.ResponseWriter(), loginUser)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=loginIn")
}

func (c *IndexController) GetLogout() {
	comm.SetLoginUser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=loginOut")
}

func (c *IndexController) GetMyprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// 验证登录
	loginUser := comm.GetLoginUser(c.Ctx.Request())
	if loginUser == nil || loginUser.Uid < 1 {
		rs["code"] = 101
		rs["msg"] = "请先登录"
		return rs
	}

	// 只读取出来最新的100次中奖记录
	list := c.ServiceResult.SearchByUser(loginUser.Uid, 1, 100)
	rs["prize_list"] = list
	// 今天抽奖次数
	day, _ := strconv.Atoi(comm.FormatFromUnixTimeShort(time.Now().Unix()))
	num := c.ServiceUserDay.Count(loginUser.Uid, day)
	rs["prize_num"] = conf.UserPrizeMax - num
	return rs
}

func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	gifts := c.ServiceGift.GetAll(true)
	giftIds := make([]int, 0)
	for _, data := range gifts {
		// 虚拟券或者实物奖才需要放到外部榜单中展示
		if data.Gtype > 1 {
			giftIds = append(giftIds, data.Id)
		}
	}
	list := c.ServiceResult.GetNewPrize(50, giftIds)
	rs["prize_list"] = list
	return rs
}
