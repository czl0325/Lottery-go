package dao

import (
	"Lottery-go/comm"
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
	"log"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine: engine,
	}
}

func (d *GiftDao) Get(id int) *models.Gift {
	data := &models.Gift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *GiftDao) GetAll() []models.Gift {
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

func (d *GiftDao) Create(data *models.Gift) error {
	_, err := d.engine.Insert(data)
	return err
}

// 获取到当前可以获取的奖品列表
// 有奖品限定，状态正常，时间期间内
// gtype倒序， displayorder正序
func (d *GiftDao) GetAllUse() []models.Gift {
	now := comm.NowUnix()
	dataList := make([]models.Gift, 0)
	err := d.engine.
		Cols("id", "title", "prize_num", "left_num", "prize_code",
			"prize_time", "img", "display_order", "gtype", "gdata").
		Desc("gtype").
		Asc("display_order").
		Where("prize_num>=?", 0).    // 有限定的奖品
		Where("sys_status=?", 0).    // 有效的奖品
		Where("time_begin<=?", now). // 时间期内
		Where("time_end>=?", now).   // 时间期内
		Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *GiftDao) IncrLeftNum(id, num int) (int64, error) {
	r, err := d.engine.Id(id).
		Incr("left_num", num).
		Update(&models.Gift{Id: id})
	return r, err
}

func (d *GiftDao) DecrLeftNum(id, num int) (int64, error) {
	r, err := d.engine.Id(id).
		Decr("left_num", num).
		Where("left_num>=?", num).
		Update(&models.Gift{Id: id})
	return r, err
}
