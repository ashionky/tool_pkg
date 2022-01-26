/**
 * @Author pibing
 * @create 2022/1/24 4:03 PM
 */

package redis

import (
	"context"
	"testing"
)

func TestRedis(t *testing.T) {
	var ctx = context.Background()
	cfg := &Conf{
		Address:     "127.0.0.1:6379",
		MaxPoolSize: 10,
		Password:    "",
		DB:          1,
		IdleTimeout: 10,
	}
	client := NewClient(cfg)
	client.Set(ctx, "b", "aaaa", 0)

}
