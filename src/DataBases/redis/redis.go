package redis

import (
	"time"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var redisPool *redis.Pool
var Conn redis.Conn

func init() {
	viper.SetConfigName("config")
    viper.AddConfigPath("config")
	viper.SetConfigType("json")
	
	err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("config file error: %s\n", err)
	}
	
	server := viper.GetString(`redis.server`)
	password := viper.GetString(`redis.password`)
	db := viper.GetInt(`redis.database`)

	redisPool = newPool(server, password, db)
	Conn = redisPool.Get()

}

// newPool New redis pool.
func newPool(server, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server,
				redis.DialPassword(password),
				redis.DialDatabase(db),
				redis.DialConnectTimeout(500*time.Millisecond),
				redis.DialReadTimeout(500*time.Millisecond),
				redis.DialWriteTimeout(500*time.Millisecond))
			if err != nil {
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