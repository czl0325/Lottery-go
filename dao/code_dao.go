package dao

import (
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao  {
	return &CodeDao{engine:engine}
}

func (d* CodeDao) Get (id int) *models.Code {
	data := models.Code{Id: id}
	ok , err := d.engine.Get(data)
	if ok && err == nil {
		return &data
	} else {
		return nil
	}
}

func (d *CodeDao) GetAll(page, size int) []models.Code {
	offset := (page - 1) * size
	dataList := make([]models.Code, 0)
	err := d.engine.Desc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Code{})
	if err != nil {
		return num
	} else {
		return 0
	}
}

func (d *CodeDao) Create (data *models.Code) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *CodeDao) Update (data *models.Code, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}