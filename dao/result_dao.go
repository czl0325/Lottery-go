package dao

import (
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{engine: engine}
}

func (d *ResultDao) Get(id int) *models.Result {
	data := &models.Result{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (d *ResultDao) GetAll(page, size int) []models.Result {
	offset := (page - 1) * size
	dataList := make([]models.Result, 0)
	err := d.engine.Desc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *ResultDao) CountAll() int64 {
	count, err := d.engine.Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *ResultDao) GetNewPrize(size int, giftIds []int) []models.Result {
	dataList := make([]models.Result, 0)
	err := d.engine.In("gift_id", giftIds).Desc("id").Limit(size).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *ResultDao) SearchByGift(giftId, page, size int) []models.Result {
	offset := (page - 1) * size
	dataList := make([]models.Result, 0)
	err := d.engine.Where("gift_id=?", giftId).Desc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *ResultDao) SearchByUser(uid, page, size int) []models.Result {
	offset := (page - 1) * size
	dataList := make([]models.Result, 0)
	err := d.engine.Where("uid=?", uid).Desc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *ResultDao) CountByGift(giftId int) int64 {
	count, err := d.engine.Where("gift_id=?", giftId).Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *ResultDao) CountByUser(uid int) int64 {
	count, err := d.engine.Where("uid=?", uid).Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *ResultDao) Delete(id int) error {
	data := &models.Result{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *ResultDao) Update(data *models.Result, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *ResultDao) Create(result *models.Result) error {
	_, err := d.engine.Insert(result)
	return err
}
