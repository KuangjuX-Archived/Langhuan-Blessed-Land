package Controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/json"
	redis "github.com/KuangjuX/Lang-Huan-Blessed-Land/Services/RedisService"
	ws "github.com/KuangjuX/Lang-Huan-Blessed-Land/Services/WebSocketService"
)

func Chat(c *gin.Context){
	hub := ws.NewHub()
	go hub.Run()
	ws.ServeWs(hub, c)
}

func GetChatInfo(c *gin.Context)  {
	room_id := c.Query("room_id")
	len, err := redis.GetListLen(room_id)
	if err != nil{
		json.JsonError(c, err)
		return
	}
	
	var end int64 = len - 1
	if len > 500 {
		end = 500
	}

	res, err := redis.GetInfoFromList(room_id, 0, end)
	if err != nil{
		json.JsonError(c, err)
		return
	}

	json.JsonDataWithSuccess(c, res)

}