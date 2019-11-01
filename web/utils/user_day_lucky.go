/**
 * 同一个User抽奖，每天的操作限制，本地或者redis缓存
 */
package utils

import (
	"Lottery-go/comm"
	"Lottery-go/datasource"
	"fmt"
	"log"
	"math"
	"time"
)

const userFrameSize = 2

func init() {
	// User当天的统计数，整点归零，设置定时器
	//duration := comm.NextDayDuration()
	//time.AfterFunc(duration, resetGroupUserList)

	// TODO: 本地开发测试的时候，每次启动归零
	resetGroupUserList()
}

// 集群模式，重置用户今天次数
func resetGroupUserList() {
	log.Println("重置用户今天次数 开始")
	cacheObj := datasource.InstanceCache()
	for i := 0; i < userFrameSize; i++ {
		key := fmt.Sprintf("day_users_%d", i)
		cacheObj.Do("DEL", key)
	}
	log.Println("重置用户今天次数 结束")
	// IP当天的统计数，整点归零，设置定时器
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupUserList)
}

// 今天的用户抽奖次数递增，返回递增后的数值
func IncrUserLuckyNum(uid int) int64 {
	i := uid % userFrameSize
	// 集群的redis统计数递增
	return incrServerUserLucyNum(i, uid)
}

func incrServerUserLucyNum(i, uid int) int64 {
	key := fmt.Sprintf("day_users_%d", i)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, uid, 1)
	if err != nil {
		log.Println("user_day_lucky redis HINCRBY key=", key,
			", uid=", uid, ", err=", err)
		return math.MaxInt32
	} else {
		num := rs.(int64)
		return num
	}
}

// 从给定的数据直接初始化用户的参与次数
func InitUserLuckyNum(uid int, num int64) {
	if num <= 1 {
		return
	}
	i := uid % userFrameSize
	// 集群
	initServUserLuckyNum(i, uid, num)
}

func initServUserLuckyNum(i, uid int, num int64) {
	key := fmt.Sprintf("day_users_%d", i)
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, uid, num)
	if err != nil {
		log.Println("user_day_lucky redis HSET key=", key,
			", uid=", uid, ", err=", err)
	}
}
