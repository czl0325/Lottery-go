package datasource

import (
	"Lottery-go/conf"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

var rdsLock sync.Mutex
var cacheInstance *RedisConn

// 封装成一个redis资源池
type RedisConn struct {
	pool 			*redis.Pool
	showDebug 		bool
}

func (rds *RedisConn) Do (commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := rds.pool.Get()
	defer conn.Close()

	t1 := time.Now().UnixNano()
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Fatal("redis请求错误,", e)
		}
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		ct.Foreground(ct.Cyan, true)
		fmt.Printf("[redis] [info] [%dus]cmd=%s, err=%s, args=%v, reply=%s\n", (t2-t1)/1000, commandName, err, args, reply)
		ct.ResetColor()
	}
	return reply, err
}

func (rds *RedisConn) ShowDebug (show bool)  {
	rds.showDebug = show
}

// 得到唯一的redis缓存实例
func InstanceCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()

	if cacheInstance != nil {
		return cacheInstance
	}
	return NewCache()
}

// 重新实例化
func NewCache() *RedisConn {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			conn, e = redis.Dial("tcp", fmt.Sprintf("%s:%d",conf.RdsCache.Host,conf.RdsCache.Port))
			if e != nil {
				log.Fatal("创建redis失败,", e)
				return nil, e
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle: 10000,
		MaxActive: 10000,
		IdleTimeout: 0,
		Wait:false,
		MaxConnLifetime:0,
	}
	instance := &RedisConn{
		pool:      &pool,
		showDebug: true,
	}
	cacheInstance = instance
	return cacheInstance
}