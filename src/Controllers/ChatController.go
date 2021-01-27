package Controllers

import(
	redis "github.com/KuangjuX/Lang-Huan-Blessed-Land/Services/RedisService"
	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
)

func GetChatInfo(c *gin.Context)  {
	room_id := c.Query("room_id")
	len, err := redis.GetListLen(room_id)
	if err != nil{
		Help.JsonError(c, err)
		return
	}

	res, err := redis.GetInfoFromList(room_id, 0, len-1)
	if err != nil{
		Help.JsonError(c, err)
		return
	}

	Help.JsonDataWithSuccess(c, res)

}