package initConn

import "github.com/go-redis/redis"

var Rdb *redis.Client

func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err = Rdb.Ping().Result()
	return err
}
