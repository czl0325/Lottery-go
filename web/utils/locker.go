/**
 *
 * 抽奖中用到的锁
 */
package utils

import (
	"Lottery-go/datasource"
	"fmt"
	)

// 加锁，抽奖的时候需要用到的锁，避免一个用户并发多次抽奖
func LockLucky(uid int) bool {
		return lockLuckyServ(uid)
}

// 解锁，抽奖的时候需要用到的锁，避免一个用户并发多次抽奖
func UnlockLucky(uid int) bool {
		return unlockLuckyServ(uid)
}

func getLuckyLockKey(uid int) string {
	return fmt.Sprintf("lucky_lock_%d", uid)
}

func lockLuckyServ(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	} else {
		return false
	}
}

func unlockLuckyServ(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("DEL", key)
	if rs == "OK" {
		return true
	} else {
		return false
	}
}

