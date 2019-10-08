package services

import (
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var cachedUserList = make(map[int]*models.User)
var cachedUserLock = sync.Mutex{}

type userService struct {
	dao *dao.UserDao
}

type UserService interface {
	GetAll(page, size int) []models.User
	CountAll() int64
	Get(id int) *models.User
	Update(user *models.User, columns []string) error
	Create(user *models.User) error
}

func NewUserService() UserService {
	return &userService{dao:dao.NewUserDao(datasource.InstanceDbMaster()),}
}

func (u userService) GetAll(page, size int) []models.User {
	return u.dao.GetAll(page, size)
}

func (u userService) CountAll() int64 {
	return u.dao.CountAll()
}

func (u userService) Get(id int) *models.User {
	return u.dao.Get(id)
}

func (u userService) Update(user *models.User, columns []string) error {
	panic("implement me")
}

func (u userService) Create(user *models.User) error {
	panic("implement me")
}

func (u *userService) GetByCache(id int) *models.User {
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
}