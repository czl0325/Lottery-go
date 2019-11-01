package controllers

import "Lottery-go/comm"

func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// 1 验证登录用户
	loginUser := comm.GetLoginUser(c.Ctx.Request())
	if loginUser == nil || loginUser.Uid <= 0 {
		rs["code"] = comm.NoLogin
		rs["msg"] = "请先登录，再来抽奖"
		return rs
	}
	//ip := comm.ClientIP(c.Ctx.Request())
	return rs
}
