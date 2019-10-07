package services

import (
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
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
	panic("implement me")
}

func (g giftService) CountAll() int64 {
	panic("implement me")
}

func (g giftService) Get(id int, useCache bool) *models.Gift {
	panic("implement me")
}

func (g giftService) Delete(id int) error {
	panic("implement me")
}

func (g giftService) Update(data *models.Gift, columns []string) error {
	panic("implement me")
}

func (g giftService) Create(data *models.Gift) error {
	panic("implement me")
}

func (g giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	panic("implement me")
}

func (g giftService) IncrLeftNum(id, num int) (int64, error) {
	panic("implement me")
}

func (g giftService) DecrLeftNum(id, num int) (int64, error) {
	panic("implement me")
}