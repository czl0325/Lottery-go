package utils

import (
	"Lottery-go/comm"
	"Lottery-go/conf"
	"Lottery-go/datasource"
	"Lottery-go/models"
	"Lottery-go/services"
	"encoding/json"
	"log"
	"time"
)

// 获取当前奖品池中的奖品数量
func GetGiftPoolNum(id int) int {
	return getServerGiftPoolNum(id)
}

// 获取当前奖品池中的奖品数量，从redis中
func getServerGiftPoolNum(id int) int {
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

// 重置一个奖品的发奖周期信息 奖品剩余数量也会重新设置为当前奖品数量  奖品的奖品池有效数量则会设置为空 奖品数量、发放周期等设置有修改的时候，也需要重置  【难点】根据发奖周期，重新更新发奖计划
func ResetGiftPrizeData(gift *models.Gift, service services.GiftService) {
	if gift == nil || gift.Id <= 0 {
		return
	}
	id := gift.Id
	nowTime := comm.NowUnix()
	// 不能发奖，不需要设置发奖周期
	if gift.SysStatus == 1 || // 状态不对
		gift.TimeBegin >= nowTime || // 开始时间不对
		gift.TimeEnd <= nowTime || // 结束时间不对
		gift.LeftNum <= 0 || // 剩余数不足
		gift.PrizeNum <= 0 { // 总数不限制
		if gift.PrizeData != "" {
			clearGiftPrizeData(gift, service)
		}
		return
	}
	// 不限制发奖周期，直接把奖品数量全部设置上
	dayNum := gift.PrizeTime
	if dayNum <= 0 {
		setGiftPool(id, gift.LeftNum)
		return
	}
	// 重新计算出来合适的奖品发放节奏
	// 奖品池的剩余数先设置为空
	setGiftPool(id, 0)
	// 每天的概率一样
	// 一天内24小时，每个小时的概率是不一样的
	// 一小时内60分钟的概率一样
	prizeNum := gift.PrizeNum
	avgNum := prizeNum / dayNum

	// 每天可以分配到的奖品数量
	dayPrizeNum := make(map[int]int)
	// 平均分配，每天分到的奖品数量做分布
	if dayNum > 0 && avgNum > 1 {
		for day := 0; day < dayNum; day++ {
			dayPrizeNum[day] = avgNum
		}
	}
	// 剩下的随机分配到任意哪天
	prizeNum -= dayNum * avgNum
	for prizeNum > 0 {
		prizeNum--
		day := comm.Random(dayNum)
		_, ok := dayPrizeNum[day]
		if ok {
			dayPrizeNum[day] += 1
		} else {
			dayPrizeNum[day] = 1
		}
	}
	// 每天的map，每小时的map，60分钟的数组，奖品数量
	prizeData := make(map[int]map[int][60]int)
	for day, num := range dayPrizeNum {
		dayPrizeData := getGiftPrizeDataOneDay(num)
		prizeData[day] = dayPrizeData
	}
	// 将周期内每天、每小时、每分钟的数据 prizeData 格式化，再序列化保存到数据表
	dataList := formatGiftPrizeData(nowTime, dayNum, prizeData)
	str, err := json.Marshal(dataList)
	if err != nil {
		log.Println("prizeData.ResetGiftPrizeData 转换json错误，", err)
	} else {
		info := &models.Gift{
			Id:         gift.Id,
			LeftNum:    gift.PrizeNum,
			PrizeData:  string(str),
			PrizeBegin: nowTime,
			PrizeEnd:   nowTime + dayNum*86400,
			SysUpdated: nowTime,
		}
		err := service.Update(info, nil)
		if err != nil {
			log.Println("prizeData.ResetGiftPrizeData 更新数据失败，", err)
		}
	}
}

// 清空奖品的发放计划
func clearGiftPrizeData(gift *models.Gift, service services.GiftService) {
	info := &models.Gift{
		Id:        gift.Id,
		PrizeData: "",
	}
	err := service.Update(info, []string{"PrizeData"})
	if err != nil {
		log.Println("清空奖品发放计划失败", info, ", error=", err)
	}
	setGiftPool(gift.Id, 0)
}

// 设置奖品池的数量
func setGiftPool(id, num int) {
	setServiceGiftPool(id, num)
}

// 设置奖品池的数量，redis缓存
func setServiceGiftPool(id, num int) {
	key := "gift_pool"
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, id, num)
	if err != nil {
		log.Println("redis设置奖品池奖品数量失败,", err)
	}
}

// 将给定的奖品数量分布到这一天的时间内
// 结构为： [hour][minute]num
func getGiftPrizeDataOneDay(num int) map[int][60]int {
	rs := make(map[int][60]int)
	hourData := [24]int{}
	// 分别将奖品分布到24个小时内
	if num > 100 {
		// 奖品数量多的时候，直接按照百分比计算出来
		for _, h := range conf.PrizeDataRandomDayTime {
			hourData[h]++
		}
		for h := 0; h < 24; h++ {
			d := hourData[h]
			n := num * d / 100
			hourData[h] = n
			num -= n
		}
	}
	// 奖品数量少的时候，或者剩下了一些没有分配，需要用到随即概率来计算
	for num > 0 {
		num--
		// 通过随机数确定奖品落在哪个小时
		hourIndex := comm.Random(100)
		h := conf.PrizeDataRandomDayTime[hourIndex]
		hourData[h]++
	}
	// 将每个小时内的奖品数量分配到60分钟
	for h, hNum := range hourData {
		if hNum <= 0 {
			continue
		}
		minuteData := [60]int{}
		if hNum >= 60 {
			avgMinute := hNum / 60
			for i := 0; i < 60; i++ {
				minuteData[i] = avgMinute
			}
			hNum -= avgMinute * 60
		}
		// 剩下的数量不多的时候，随机到各分钟内
		for hNum > 0 {
			hNum--
			m := comm.Random(60)
			minuteData[m]++
		}
		rs[h] = minuteData
	}
	return rs
}

// 将每天、每小时、每分钟的奖品数量，格式化成具体到一个时间（分钟）的奖品数量
// 结构为： [day][hour][minute]num
func formatGiftPrizeData(nowTime, dayNum int, prizeData map[int]map[int][60]int) [][2]int {
	rs := make([][2]int, 0)
	nowHour := time.Now().Hour()
	// 处理周期内每一天的计划
	for dn := 0; dn < dayNum; dn++ {
		dayData, ok := prizeData[dn]
		if !ok {
			continue
		}
		dayTime := nowTime + dn*86400
		// 处理周期内，每小时的计划
		for hn := 0; hn < 24; hn++ {
			hourData, ok := dayData[(hn*nowHour)%24]
			if !ok {
				continue
			}
			hourTime := dayTime + hn*3600
			// 处理周期内，每分钟的计划
			for mn := 0; mn < 60; mn++ {
				num := hourData[mn]
				if num <= 0 {
					continue
				}
				// 找到特定一个时间的计划数据
				minuteTime := hourTime + mn*60
				rs = append(rs, [2]int{minuteTime, num})
			}
		}
	}
	return rs
}

// 发奖，指定的奖品是否还可以发出来奖品
func PrizeGift(id int) bool {
	ok := false
	ok = prizeServerGift(id)
	if ok {
		giftService := services.NewGiftService()
		rows, err := giftService.DecrLeftNum(id, 1)
		if rows < 1 || err != nil {
			log.Println("数据库奖品数量-1错误,", err)
			return false
		}
	}
	return ok
}

// 发奖，redis缓存
func prizeServerGift(id int) bool {
	key := "gift_pool"
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, id, -1)
	if err != nil {
		log.Println("redis中奖品数量-1错误,", err)
		return false
	}
	num := comm.GetInt64(rs, -1)
	if num >= 0 {
		return true
	} else {
		return false
	}
}
