package models

type Code struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	GiftId     int    `xorm:"comment('奖品ID，关联gift表') INT(13)"`
	Code       string `xorm:"comment('虚拟券编码') VARCHAR(255)"`
	SysCreated int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated int    `xorm:"comment('更新时间') INT(13)"`
	SysStatus  int    `xorm:"comment('0正常，1作废，2已发放') INT(11)"`
}
