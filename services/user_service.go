package services

import (
	"Lottery-go/comm"
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
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
	user := u.GetByCache(id)
	if user == nil || user.Id <= 0 {
		user = u.dao.Get(id)
		if user == nil || user.Id <= 0 {
			user = &models.User{Id: id}
		}
		u.SetByCache(user)
	}
	return u.dao.Get(id)
}

func (u userService) Update(user *models.User, columns []string) error {
	u.UpdateByCache(user, columns)
	return u.dao.Update(user, columns)
}

func (u userService) Create(user *models.User) error {
	return u.dao.Create(user)
}

func (u *userService) GetByCache(id int) *models.User {
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		log.Println("user_service.getByCache HGETALL key=", key, ", error=", err)
		return nil
	}
	dataId := comm.GetInt64FromStringMap(dataMap, "Id", 0)
	if dataId <= 0 {
		return nil
	}
	data := &models.User{
		Id:         int(dataId),
		UserName:   comm.GetStringFromStringMap(dataMap, "UserName", ""),
		BlackTime:  int(comm.GetInt64FromStringMap(dataMap, "BlackTime", 0)),
		RealName:   comm.GetStringFromStringMap(dataMap, "RealName", ""),
		Mobile:     comm.GetStringFromStringMap(dataMap, "Mobile", ""),
		Address:    comm.GetStringFromStringMap(dataMap, "Address", ""),
		SysCreated: int(comm.GetInt64FromStringMap(dataMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(dataMap, "SysUpdated", 0)),
		SysIp:      comm.GetStringFromStringMap(dataMap, "SysIp", ""),
	}
	return data
}

func (u *userService) SetByCache (data *models.User) {
	if data == nil || data.Id <= 0 {
		return
	}
	id := data.Id
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	// 数据更新到redis缓存
	params := []interface{}{key}
	params = append(params, "Id", id)
	params = append(params, "UserName", data.UserName)
	params = append(params, "BlackTime", data.BlackTime)
	params = append(params, "RealName", data.RealName)
	params = append(params, "Mobile", data.Mobile)
	params = append(params, "Address", data.Address)
	params = append(params, "SysCreated", data.SysCreated)
	params = append(params, "SysUpdated", data.SysUpdated)
	params = append(params, "SysIp", data.SysIp)
	_, err := rds.Do("HMSET", params)
	if err != nil {
		log.Println("user_service.setByCache HMSET params=", params, ", error=", err)
	}
}

func (u *userService) UpdateByCache (data *models.User, columns []string)  {
	if data == nil || data.Id <= 0 {
		return
	}
	key := fmt.Sprintf("info_user_%d", data.Id)
	rds := datasource.InstanceCache()
	// 删除redis中的缓存
	rds.Do("DEL", key)
}