/**
 * @Author pibing
 * @create 2022/1/24 5:00 PM
 */

package mysql

import (
	"fmt"
	"time"
)

type Conf struct {
	Address     string
	MaxOpenConn []int
	MaxIdleConn []int
	MaxLifeTime []time.Duration
}

func NewMysql(cfg *Conf) *DB {
	mysqlDB, err := Open("mysql", cfg.Address)
	if err != nil {
		panic(err)
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	mysqlDB.SetMaxOpenConnections(cfg.MaxOpenConn)
	for i, v := range cfg.MaxLifeTime {
		cfg.MaxLifeTime[i] = v * time.Second
	}
	mysqlDB.SetConnMaxLifetime(cfg.MaxLifeTime)
	fmt.Println("Mysql is connected!")
	return mysqlDB
}
