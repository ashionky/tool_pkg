/**
 * @Author pibing
 * @create 2022/1/24 5:16 PM
 */

package mysql

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {
	var ctx = context.Background()
	cfg := &Conf{
		Address:     "root:123456@/test?charset=utf8&parseTime=True&loc=Local;root:123456@/test?charset=utf8&parseTime=True&loc=Local",
		MaxOpenConn: []int{10, 10},
		MaxIdleConn: []int{10, 10},
		MaxLifeTime: []time.Duration{8, 8},
	}
	client := NewMysql(cfg)
	count := 0
	err := client.GetContext(ctx, &count, "select count(*) from test")
	fmt.Println(err, count)

}
