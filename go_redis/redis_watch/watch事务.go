package redis_watch

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//在用户使用事务管道的时候 先使用watch监控某个值的变化 在执行事务的exec的时候会看watch监控的值是否被修改过
//修改过的话 事务回滚不执行 反之执行
// watch监听某个key值的变化

func Watch1(rdb *redis.Client) {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rdb.Get(key).Val())
}
