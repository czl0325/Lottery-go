package models

type Blackip struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Ip         string `xorm:"comment('IP地址') VARCHAR(50)"`
	BlackTime  int    `xorm:"comment('黑名单限制到期时间') INT(11)"`
	SysCreated int    `xorm:"comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"comment('修改时间') INT(11)"`
}
