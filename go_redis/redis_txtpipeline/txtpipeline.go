package redis_txtpipeline

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 这个管道是具有事务性的 可以将执行的多条命令进行包裹
// 可以理解位锁
func TxPieline(rdb *redis.Client) {
	txpiple := rdb.TxPipeline()
	incr := txpiple.Incr("ceshi_counter")
	txpiple.Expire("ceshi_counter", time.Hour)

	_, err := txpiple.Exec()
	fmt.Println(incr.Val(), err)
}
