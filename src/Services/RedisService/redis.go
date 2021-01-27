package RedisService

import(
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/redis"
)

func GetInfoFromList(key string, start, end int64)([]string, error){
	redisConn := redis.RedisPool.Get()
	query, err := redisConn.Do("LRANGE", key, start, end)
	if err != nil{
		return nil, err
	}
	var res []string
	res, err = redis.Strings(query)
	if err != nil{
		return nil, err
	}

	return res, err
}

func GetListLen(key string)(int64, error){
	redisConn := redis.RedisPool.Get()
	res, err := redisConn.Do("llen", key)
	if err != nil{
		return 0, err
	}
	return res.(int64), err
}
