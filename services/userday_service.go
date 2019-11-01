package services

import (
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
	"fmt"
	"strconv"
	"time"
)

type userdayService struct {
	dao *dao.UserDayDao
}

type UserdayService interface {
	GetAll(page, size int) []models.Userday
	CountAll() int64
	Search(uid, day int) []models.Userday
	Count(uid, day int) int
	Get(id int) *models.Userday
	Update(user *models.Userday, columns []string) error
	Create(user *models.Userday) error
	GetUserToday(uid int) *models.Userday
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao: dao.NewUserDayDao(datasource.InstanceDbMaster()),
	}
}

func (u userdayService) GetAll(page, size int) []models.Userday {
	return u.dao.GetAll(page, size)
}

func (u userdayService) CountAll() int64 {
	return u.dao.CountAll()
}

func (u userdayService) Search(uid, day int) []models.Userday {
	return u.dao.Search(uid, day)
}

func (u userdayService) Count(uid, day int) int {
	return u.dao.Count(uid, day)
}

func (u userdayService) Get(id int) *models.Userday {
	return u.dao.Get(id)
}

func (u userdayService) Update(user *models.Userday, columns []string) error {
	return u.dao.Update(user, columns)
}

func (u userdayService) Create(user *models.Userday) error {
	return u.dao.Create(user)
}

func (u userdayService) GetUserToday(uid int) *models.Userday {
	y, m, d := time.Now().Date()
	strDate := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDate)
	list := u.dao.Search(uid, day)
	if list != nil && len(list) > 0 {
		return &list[0]
	} else {
		return nil
	}
}
