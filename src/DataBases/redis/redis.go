package redis

import (
	"time"
	"fmt"

	"github.com/garyburd/redigo/redis"
	. "github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/config"
)

var RedisPool *redis.Pool
type RedisConn = redis.Conn


func init() {

	server := REDIS.Server
	password := REDIS.Password
	db := REDIS.Database

	RedisPool = newPool(server, password, db)
	

}

// newPool New redis pool.
func newPool(server, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 256,
		MaxActive: 0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server,
				redis.DialPassword(password),
				redis.DialDatabase(db),
				redis.DialConnectTimeout(500*time.Millisecond),
				redis.DialReadTimeout(500*time.Millisecond),
				redis.DialWriteTimeout(500*time.Millisecond))
			if err != nil {
				fmt.Printf("error: %s\n", err)
				return nil, err
			}

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < 5*time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func Strings(reply interface{})([]string, error){
	var err error
	res, err := redis.Strings(reply, err)
	if err != nil{
		return nil, err
	}
	return res, nil
}