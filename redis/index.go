/**
 * @Author pibing
 * @create 2022/1/24 4:02 PM
 */

package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Conf struct {
	Address     string
	Username    string
	Password    string
	DB          int
	MaxPoolSize int
	IdleTimeout int
}

func NewClient(conf *Conf) *redis.Client {

	redisOptions := &redis.Options{
		Addr:        conf.Address,
		Username:    conf.Username,
		Password:    conf.Password,
		DB:          conf.DB,
		PoolSize:    conf.MaxPoolSize,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
	}

	client := redis.NewClient(redisOptions)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis is connected!")
	return client
}
