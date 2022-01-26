/**
 * @Author pibing
 * @create 2022/1/24 4:25 PM
 */

package mongo

import (
	"context"
	"fmt"
	"testing"
)

func TestMongo(t *testing.T) {
	var ctx = context.Background()
	cfg := &Conf{
		Address:       "mongodb://root:123456@127.0.0.1:27017",
		MaxPoolSize:   10,
		SocketTimeout: 5,
	}
	NewClient(cfg)
	count, err := MongoDBCurd("test", "test").Count(ctx, B{})
	fmt.Println(count, err)
}
