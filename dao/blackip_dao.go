package dao

import (
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
)

type BlackIpDao struct {
	engine *xorm.Engine
}

func NewBlackIpDao(engine *xorm.Engine) *BlackIpDao  {
	return &BlackIpDao{engine:engine}
}

func (d *BlackIpDao) Get(id int) *models.Blackip  {
	data := &models.Blackip{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (d *BlackIpDao) GetAll(page, size int) []models.Blackip {
	offset := (page - 1) * size
	dataList := make([]models.Blackip, 0)
	err := d.engine.Desc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d* BlackIpDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Blackip{})
	if err != nil {
		return 0
	}
	return num
}

func (d* BlackIpDao) Search(ip string) []models.Blackip {
	dataList := make([]models.Blackip, 0)
	err := d.engine.Where("ip=?", ip).Desc("id").Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *BlackIpDao) Create (data *models.Blackip) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackIpDao) Update (data *models.Blackip, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}