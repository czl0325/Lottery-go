package services

import (
	"Lottery-go/comm"
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type GiftService interface {
	GetAll(useCache bool) []models.Gift
	CountAll() int64
	//Search(country string) []models.LtGift
	Get(id int, useCache bool) *models.Gift
	Delete(id int) error
	Update(data *models.Gift, columns []string) error
	Create(data *models.Gift) error
	GetAllUse(useCache bool) []models.ObjGiftPrize
	IncrLeftNum(id, num int) (int64, error)
	DecrLeftNum(id, num int) (int64, error)
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{dao:dao.NewGiftDao(datasource.InstanceDbMaster()),}
}

func (g giftService) GetAll(useCache bool) []models.Gift {
	if !useCache {
		return g.dao.GetAll()
	}
	// 先读取缓存
	gifts := g.getAllByCache()
	if len(gifts) < 1 {
		gifts = g.dao.GetAll()
		g.setAllByCache(gifts)
	}
	return gifts
}

func (g giftService) CountAll() int64 {
	return g.dao.CountAll()
}

func (g giftService) Get(id int, useCache bool) *models.Gift {
	if !useCache {
		return g.dao.Get(id)
	}
	gifts := g.getAllByCache()
	for _, gift := range gifts {
		if gift.Id == id {
			return &gift
		}
	}
	return nil
}

func (g giftService) Delete(id int) error {
	g.cleanCache()
	return g.dao.Delete(id)
}

func (g giftService) Update(data *models.Gift, columns []string) error {
	g.cleanCache()
	return g.dao.Update(data, columns)
}

func (g giftService) Create(data *models.Gift) error {
	g.cleanCache()
	return g.dao.Create(data)
}

func (g giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	list := make([]models.Gift, 0)
	if !useCache {
		list = g.dao.GetAllUse()
	} else {
		now := comm.NowUnix()
		gifts := g.GetAll(true)
		for _, gift := range gifts {
			if gift.Id > 0 && gift.SysStatus == 0 && gift.PrizeNum >= 0 && gift.TimeBegin <= now && gift.TimeEnd >= now {
				list = append(list, gift)
			}
		}
	}

	if list != nil {
		gifts := make([]models.ObjGiftPrize, 0)
		for _, gift := range list {
			codes := strings.Split(gift.PrizeCode, "-")
			if len(codes) == 2 {
				codeA := codes[0]
				codeB := codes[1]
				a, e1 := strconv.Atoi(codeA)
				b, e2 := strconv.Atoi(codeB)
				if e1 == nil && e2 == nil && b >= a && a > 0 && b < 10000 {
					data := models.ObjGiftPrize{
						Id:           gift.Id,
						Title:        gift.Title,
						PrizeNum:     gift.PrizeNum,
						LeftNum:      gift.LeftNum,
						PrizeCodeA:   a,
						PrizeCodeB:   b,
						Img:          gift.Img,
						Displayorder: gift.DisplayOrder,
						Gtype:        gift.Gtype,
						Gdata:        gift.Gdata,
					}
					gifts = append(gifts, data)
				}
			}
		}
		return gifts
	}
	return []models.ObjGiftPrize{}
}

func (g giftService) IncrLeftNum(id, num int) (int64, error) {
	return g.dao.IncrLeftNum(id, num)
}

func (g giftService) DecrLeftNum(id, num int) (int64, error) {
	return g.dao.DecrLeftNum(id, num)
}

// 从缓存中获取全部的奖品
func (s* giftService) getAllByCache() []models.Gift {
	key := "allgift"
	rds := datasource.InstanceCache()
	rs, err := rds.Do("GET", key)
	if err != nil {
		log.Println("gift_service.getAllByCache GET key=", key, ", error=", err)
		return nil
	}
	str := comm.GetString(rs, "")
	if str == "" {
		return nil
	}
	// 将json数据反序列化
	dataList := []map[string]interface{}{}
	err = json.Unmarshal([]byte(str), &dataList)
	if err != nil {
		log.Println("gift_service.getAllByCache json.Unmarshal error=", err)
		return nil
	}
	// 格式转换
	gifts := make([]models.Gift, len(dataList))
	for i := 0; i < len(dataList); i++ {
		data := dataList[i]
		id := comm.GetInt64FromMap(data, "Id", 0)
		if id <= 0 {
			gifts[i] = models.Gift{}
		} else {
			gift := models.Gift{
				Id:           int(id),
				Title:        comm.GetStringFromMap(data, "Title", ""),
				PrizeNum:     int(comm.GetInt64FromMap(data, "PrizeNum", 0)),
				LeftNum:      int(comm.GetInt64FromMap(data, "LeftNum", 0)),
				PrizeCode:    comm.GetStringFromMap(data, "PrizeCode", ""),
				PrizeTime:    int(comm.GetInt64FromMap(data, "PrizeTime", 0)),
				Img:          comm.GetStringFromMap(data, "Img", ""),
				DisplayOrder: int(comm.GetInt64FromMap(data, "DisplayOrder", 0)),
				Gtype:        int(comm.GetInt64FromMap(data, "Gtype", 0)),
				Gdata:        comm.GetStringFromMap(data, "Gdata", ""),
				TimeBegin:    int(comm.GetInt64FromMap(data, "TimeBegin", 0)),
				TimeEnd:      int(comm.GetInt64FromMap(data, "TimeEnd", 0)),
				//PrizeData:    comm.GetStringFromMap(data, "PrizeData", ""),
				PrizeBegin: int(comm.GetInt64FromMap(data, "PrizeBegin", 0)),
				PrizeEnd:   int(comm.GetInt64FromMap(data, "PrizeEnd", 0)),
				SysStatus:  int(comm.GetInt64FromMap(data, "SysStatus", 0)),
				SysCreated: int(comm.GetInt64FromMap(data, "SysCreated", 0)),
				SysUpdated: int(comm.GetInt64FromMap(data, "SysUpdated", 0)),
				SysIp:      comm.GetStringFromMap(data, "SysIp", ""),
			}
			gifts[i] = gift
		}
	}
	return gifts
}

// 将奖品的数据更新到缓存
func (s* giftService) setAllByCache(gifts []models.Gift)  {
	// 集群模式，redis缓存
	strValue := ""
	if len(gifts) > 0 {
		dataList := make([]map[string]interface{}, len(gifts))
		for i := 0; i < len(gifts); i++ {
			gift := gifts[i]
			data := make(map[string]interface{})
			data["Id"] = gift.Id
			data["Title"] = gift.Title
			data["PrizeNum"] = gift.PrizeNum
			data["LeftNum"] = gift.LeftNum
			data["PrizeCode"] = gift.PrizeCode
			data["PrizeTime"] = gift.PrizeTime
			data["Img"] = gift.Img
			data["DisplayOrder"] = gift.DisplayOrder
			data["Gtype"] = gift.Gtype
			data["Gdata"] = gift.Gdata
			data["TimeBegin"] = gift.TimeBegin
			data["TimeEnd"] = gift.TimeEnd
			//data["PrizeData"] = gift.PrizeData
			data["PrizeBegin"] = gift.PrizeBegin
			data["PrizeEnd"] = gift.PrizeEnd
			data["SysStatus"] = gift.SysStatus
			data["SysCreated"] = gift.SysCreated
			data["SysUpdated"] = gift.SysUpdated
			data["SysIp"] = gift.SysIp
			dataList[i] = data
		}
		str, err := json.Marshal(dataList)
		if err != nil {
			log.Println("转换成json失败=", err)
		}
		strValue = string(str)
	}
	key := "allgift"
	rds := datasource.InstanceCache()
	_, err := rds.Do("SET", key, strValue)
	if err != nil {
		log.Println("redis更新缓存失败, error=", err)
	}
}

// 数据更新，需要更新缓存，直接清空缓存数据
func (s *giftService) cleanCache() {
	key := "allgift"
	rds := datasource.InstanceCache()
	// 删除redis中的缓存
	rds.Do("DEL", key)
}