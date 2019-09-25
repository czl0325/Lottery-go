package models

type Userday struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Uid        int    `xorm:"comment('用户ID') INT(11)"`
	Day        string `xorm:"comment('日期，如：20180725') VARCHAR(255)"`
	Num        int    `xorm:"comment('次数') INT(11)"`
	SysCreated int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated int    `xorm:"comment('修改时间') INT(13)"`
}
