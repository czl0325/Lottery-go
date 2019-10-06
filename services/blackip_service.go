package services

import (
	"Lottery-go/dao"
	"Lottery-go/datasource"
	"Lottery-go/models"
	"sync"
)

// IP信息，可以缓存(本地或者redis)，有更新的时候，再根据具体情况更新缓存
var cachedBlackIpList = make(map[string]*models.Blackip)
var cachedBlackIpLock = sync.Mutex{}

type BlackIpService interface {
	GetAll(page, size int) []models.Blackip
	CountAll() int64
	Search(ip string) []models.Blackip
	Get(id int) *models.Blackip
	Update(blackip *models.Blackip, columns []string) error
	Create(blackip *models.Blackip) error
	GetByIp(ip string) *models.Blackip
}

type blackIpService struct {
	dao *dao.BlackIpDao
}

func NewBlackIpService() BlackIpService  {
	return &blackIpService{
		dao:dao.NewBlackIpDao(datasource.InstanceDbMaster()),
	}
}

func (b blackIpService) GetAll(page, size int) []models.Blackip {
	return b.dao.GetAll(page, size)
}

func (b blackIpService) CountAll() int64 {
	return b.dao.CountAll()
}

func (b blackIpService) Search(ip string) []models.Blackip {
	return b.dao.Search(ip)
}

func (b blackIpService) Get(id int) *models.Blackip {
	return b.dao.Get(id)
}

func (b blackIpService) Update(blackip *models.Blackip, columns []string) error {
	return b.dao.Update(blackip, columns)
}

func (b blackIpService) Create(blackip *models.Blackip) error {
	return b.dao.Create(blackip)
}

func (b blackIpService) GetByIp(ip string) *models.Blackip {
	return b.dao.GetByIp(ip)
}