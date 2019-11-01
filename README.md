# Lottery-go
做高并发的抽奖系统后台


### 本项目需要用到的库：

```
go get -u github.com/kataras/iris
go get -u github.com/go-sql-driver/mysql
go get -u github.com/go-xorm/xorm
```

### 表

用户表

```
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	UserName   string `xorm:"comment('用户名') VARCHAR(50)"`
	BlackTime  int    `xorm:"default 0 comment('黑名单限制到期时间') INT(13)"`
	RealName   string `xorm:"comment('联系人') VARCHAR(50)"`
	Mobile     string `xorm:"comment('手机号') VARCHAR(50)"`
	Address    string `xorm:"comment('联系地址') TEXT"`
	SysCreated int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated int    `xorm:"comment('修改时间') INT(13)"`
	SysIp      string `xorm:"comment('IP地址') VARCHAR(50)"`
```
奖品表
```
Id           int    `xorm:"not null pk autoincr INT(11)"`
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
	SysStatus    int    `xorm:"not null default 0 comment('状态，0 正常，1 删除') INT(13)"`
	SysCreated   int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated   int    `xorm:"comment('修改时间') INT(13)"`
	SysIp        string `xorm:"comment('操作人IP') VARCHAR(50)"`
```

优惠券表
```
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	GiftId     int    `xorm:"comment('奖品ID，关联gift表') INT(13)"`
	Code       string `xorm:"comment('虚拟券编码') VARCHAR(255)"`
	SysCreated int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated int    `xorm:"comment('更新时间') INT(13)"`
	SysStatus  int    `xorm:"comment('0正常，1作废，2已发放') INT(11)"`
  ```
结果表

```
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
  ```

黑名单表

```
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Ip         string `xorm:"comment('IP地址') VARCHAR(50)"`
	BlackTime  int    `xorm:"comment('黑名单限制到期时间') INT(11)"`
	SysCreated int    `xorm:"comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"comment('修改时间') INT(11)"`
  ```
  
  用户每日抽奖次数表
  
```
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Uid        int    `xorm:"comment('用户ID') INT(11)"`
	Day        string `xorm:"comment('日期，如：20180725') VARCHAR(255)"`
	Num        int    `xorm:"comment('次数') INT(11)"`
	SysCreated int    `xorm:"comment('创建时间') INT(13)"`
	SysUpdated int    `xorm:"comment('修改时间') INT(13)"`
  ```
