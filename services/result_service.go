package services

import (
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
)

type ResultService interface {
	GetAll(page, size int) []models.Result
	CountAll() int64
	GetNewPrize(size int, giftIds []int) []models.Result
	SearchByGift(giftId, page, size int) []models.Result
	SearchByUser(uid, page, size int) []models.Result
	CountByGift(giftId int) int64
	CountByUser(uid int) int64
	Get(id int) *models.Result
	Delete(id int) error
	Update(user *models.Result, columns []string) error
	Create(user *models.Result) error
}

type resultService struct {
	dao *dao.ResultDao
}



func NewResultService() ResultService  {
	return &resultService{dao:dao.NewResultDao(datasource.InstanceDbMaster()),}
}

func (r resultService) GetAll(page, size int) []models.Result {
	return r.dao.GetAll(page, size)
}

func (r resultService) CountAll() int64 {
	return r.dao.CountAll()
}

func (r resultService) GetNewPrize(size int, giftIds []int) []models.Result {
	return r.dao.GetNewPrize(size, giftIds)
}

func (r resultService) SearchByGift(giftId, page, size int) []models.Result {
	return r.dao.SearchByGift(giftId, page, size)
}

func (r resultService) SearchByUser(uid, page, size int) []models.Result {
	return r.dao.SearchByUser(uid, page, size)
}

func (r resultService) CountByGift(giftId int) int64 {
	return r.dao.CountByGift(giftId)
}

func (r resultService) CountByUser(uid int) int64 {
	return r.dao.CountByUser(uid)
}

func (r resultService) Get(id int) *models.Result {
	return r.dao.Get(id)
}

func (r resultService) Delete(id int) error {
	return r.dao.Delete(id)
}

func (r resultService) Update(user *models.Result, columns []string) error {
	return r.dao.Update(user, columns)
}

func (r resultService) Create(user *models.Result) error {
	return r.dao.Create(user)
}