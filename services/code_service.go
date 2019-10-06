package services

import (
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
)

type CodeService interface {
	GetAll(page, size int) []models.Code
	CountAll() int64
	CountByGift(giftId int) int64
	Search(giftId int) []models.Code
	Get(id int) *models.Code
	Delete(id int) error
	Update(code *models.Code, columns []string) error
	Create(code *models.Code) error
	NextUsingCode(giftId, codeId int) *models.Code
	UpdateByCode(code *models.Code, columns []string) error
} 

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{dao:dao.NewCodeDao(datasource.InstanceDbMaster()),}
}

func (c codeService) GetAll(page, size int) []models.Code {
	return c.dao.GetAll(page, size)
}

func (c codeService) CountAll() int64 {
	return c.dao.CountAll()
}

func (c codeService) CountByGift(giftId int) int64 {
	return c.dao.CountByGift(giftId)
}

func (c codeService) Search(giftId int) []models.Code {
	return c.dao.Search(giftId)
}

func (c codeService) Get(id int) *models.Code {
	return c.dao.Get(id)
}

func (c codeService) Delete(id int) error {
	return c.dao.Delete(id)
}

func (c codeService) Update(code *models.Code, columns []string) error {
	return c.dao.Update(code, columns)
}

func (c codeService) Create(code *models.Code) error {
	return c.dao.Create(code)
}

func (c codeService) NextUsingCode(giftId, codeId int) *models.Code {
	return c.dao.NextUsingCode(giftId, codeId)
}

func (c codeService) UpdateByCode(code *models.Code, columns []string) error {
	return c.dao.UpdateByCode(code, columns)
}

