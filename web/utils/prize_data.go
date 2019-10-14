package utils

import (
	"Lottery-go/comm"
	"Lottery-go/datasource"
	"log"
)

// 获取当前奖品池中的奖品数量
func GetGiftPoolNum(id int) int {
	return getServerGiftPoolNum(id)
}

// 获取当前奖品池中的奖品数量，从redis中
func getServerGiftPoolNum(id int) int  {
	key := "gift_pool"
	cache := datasource.InstanceCache()
	rs, err := cache.Do("HGET", key, id)
	if err != nil {
		log.Println("从redis中获取当前奖品池的奖品数量失败,", err)
		return 0
	}
	num := comm.GetInt64(rs, 0)
	return int(num)
}