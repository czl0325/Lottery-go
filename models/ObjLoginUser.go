package models

// 站点中与浏览器交互的用户模型
type ObjLoginUser struct {
	Uid      int
	Username string
	Now      int
	Ip       string
	Sign     string
}