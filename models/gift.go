package models

type Gift struct {
	Id           string `xorm:"not null pk VARCHAR(255)"`
	Title        string `xorm:"not null comment('奖品名称') VARCHAR(255)"`
	PrizeNum     int    `xorm:"not null default -1 comment('奖品数量，0 无限量，>0限量，<0无奖品') INT(10)"`
	LeftNum      int    `xorm:"not null default 0 comment('剩余数量') INT(10)"`
	PrizeCode    string `xorm:"not null comment('0-9999表示100%，0-0表示万分之一的中奖概率') VARCHAR(255)"`
	PrizeTime    int    `xorm:"not null default 0 comment('发奖周期，D天') INT(10)"`
	Img          string `xorm:"comment('奖品图片') VARCHAR(255)"`
	DisplayOrder int    `xorm:"comment('位置序号，小的排在前面') INT(255)"`
	Gtype        int    `xorm:"not null comment('奖品类型，0 虚拟币，1 虚拟券，2 实物-小奖，3 实物-大奖'') INT(10)"`
	Gdata        string `xorm:"comment('扩展数据，如：虚拟币数量') TEXT"`
	TimeBegin    int    `xorm:"not null comment('开始时间') INT(13)"`
	TimeEnd      int    `xorm:"not null comment('结束时间') INT(13)"`
	PrizeData    string `xorm:"comment('发奖计划，[[时间1,数量1],[时间2,数量2]]') TEXT"`
	PrizeBegin   int    `xorm:"comment('发奖计划周期的开始') INT(13)"`
	PrizeEnd     int    `xorm:"comment('发奖计划周期的结束') INT(13)"`
	SysStatus    int    `xorm:"comment('状态，0 正常，1 删除') INT(13)"`
	SysCreated   int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated   int    `xorm:"comment('修改时间') INT(13)"`
	SysIp        string `xorm:"comment('操作人IP') VARCHAR(50)"`
}
