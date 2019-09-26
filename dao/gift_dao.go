package dao

import (
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
	"log"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine:engine,
	}
}

func (d *GiftDao) Get (id int) *models.Gift {
	data := &models.Gift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err != nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *GiftDao) GetAll () []models.Gift  {
	dataList := make([]models.Gift, 0)
	err := d.engine.Asc("sys_status").Asc("display_order").Find(&dataList)
	if err != nil {
		log.Fatal("查找全部错误:", err)
	}
	return dataList
}

func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Gift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *GiftDao) Delete(id int) error {
	data := &models.Gift{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *GiftDao) Update(data *models.Gift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *GiftDao) Create(data *models.Gift) error  {
	_, err := d.engine.Insert(data)
	return err
}