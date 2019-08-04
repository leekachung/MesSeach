package cache

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"fmt"
	"os"
)

// Redis连接池实例
var RedisPool *redis.Pool

// Add Connect Func
func redisConn(ip, port, passwd string) (redis.Conn, error) {
	c,err := redis.Dial("tcp",
		ip + ":" + port,
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
		redis.DialPassword(passwd),
		redis.DialKeepAlive(1*time.Second),
	)
	return c, err
}

// Add Connect Pool
func newPool(ip, port, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:            5,    //定义redis连接池中最大的空闲链接为5
		MaxActive:          18,    //在给定时间已分配的最大连接数(限制并发数)
		IdleTimeout:        240 * time.Second,
		MaxConnLifetime:    300 * time.Second,
		Dial:               func() (redis.Conn,error) { return redisConn(ip,port,passwd) },
	}
}

// Get Connect
func GetConn(pool *redis.Pool) (redis.Conn) {
	conn := pool.Get()
	return conn
}

func Redis () {
	fmt.Println(os.Getenv("REDIS_IP"))
	pool := newPool(os.Getenv("REDIS_IP"), 
	os.Getenv("REDIS_PORT"), 
	os.Getenv("REDIS_PW"))
	RedisPool = pool
}
