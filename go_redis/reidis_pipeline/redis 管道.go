package reidis_pipeline

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Ceshi(Rdb *redis.Client) {

	piple := Rdb.Pipeline()

	incr := piple.Incr("ceshi_counter")
	piple.Expire("ceshi_counter", time.Hour)

	_, err := piple.Exec()
	fmt.Println(incr.Val(), err)
}
