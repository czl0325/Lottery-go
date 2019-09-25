package models

type User struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	UserName   string `xorm:"comment('用户名') VARCHAR(50)"`
	BlackTime  int    `xorm:"default 0 comment('黑名单限制到期时间') INT(13)"`
	RealName   string `xorm:"comment('联系人') VARCHAR(50)"`
	Mobile     string `xorm:"comment('手机号') VARCHAR(50)"`
	Address    string `xorm:"comment('联系地址') TEXT"`
	SysCreated int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated int    `xorm:"comment('修改时间') INT(13)"`
	SysIp      string `xorm:"comment('IP地址') VARCHAR(50)"`
}
