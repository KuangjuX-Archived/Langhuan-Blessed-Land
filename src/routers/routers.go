package routers

import(
	"github.com/gin-gonic/gin"
    ."github.com/KuangjuX/Lang-Huan-Blessed-Land/api"
)

func InnitRouter() *gin.Engine{
    routers := gin.Default()
    routers.GET("/", Home)
    users:= routers.Group("/users")
    {
        users.GET("/",Users)
    }

    return routers;
}