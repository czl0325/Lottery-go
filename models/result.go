package models

type Result struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	GiftId     int    `xorm:"comment('奖品ID，关联gift表') INT(11)"`
	GiftName   string `xorm:"comment('奖品名称') VARCHAR(255)"`
	GiftType   int    `xorm:"comment('奖品类型，同gift. gtype') INT(11)"`
	Uid        int    `xorm:"comment('用户ID') INT(11)"`
	UserName   string `xorm:"comment('用户名') VARCHAR(50)"`
	PrizeCode  int    `xorm:"comment('抽奖编号（4位的随机数）') INT(11)"`
	GiftData   string `xorm:"comment('获奖信息') TEXT"`
	SysCreated int    `xorm:"comment('创建时间') INT(11)"`
	SysIp      string `xorm:"comment('用户抽奖的IP') VARCHAR(50)"`
	SysStatus  int    `xorm:"comment('状态，0 正常，1删除，2作弊') INT(11)"`
}
