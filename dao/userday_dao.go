package dao

import (
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
)

type UserDayDao struct {
	engine *xorm.Engine
}

func NewUserDayDao(engine *xorm.Engine) *UserDayDao {
	return &UserDayDao{engine:engine}
}

func (d *UserDayDao) Get (id int) *models.Userday {
	data := &models.Userday{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (d *UserDayDao) GetAll(page ,size int) []models.Userday {
	offset := (page - 1) * size
	dataList := make([]models.Userday, 0)
	err := d.engine.Asc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *UserDayDao) CountAll() int64 {
	count, err := d.engine.Count(&models.Userday{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *UserDayDao) Create(data* models.Userday) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *UserDayDao) Update(data* models.Userday, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}