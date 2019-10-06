package dao

import (
	"Lottery-go/models"
	"github.com/go-xorm/xorm"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{engine:engine}
}

func (d *UserDao) Get(id int) *models.User {
	data := &models.User{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (d *UserDao) GetAll(page, size int) []models.User {
	offset := (page - 1) * size
	dataList := make([]models.User, 0)
	err := d.engine.Asc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *UserDao) CountAll() int64 {
	count, err := d.engine.Count(&models.User{})
	if err != nil {
		 return 0
	} else {
		return count
	}
}

func (d *UserDao) Create (user *models.User) error {
	_, err := d.engine.Insert(user)
	return err
}

func (d *UserDao) Update (user *models.User, columns []string) error {
	_, err := d.engine.Id(user.Id).MustCols(columns...).Update(user)
	return err
}