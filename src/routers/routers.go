package routers

import(
	"github.com/gin-gonic/gin"
    ."github.com/KuangjuX/Lang-Huan-Blessed-Land/api"
)

func InnitRouter() *gin.Engine{
    routers := gin.Default()
    root := routers.Group("/")
    {   
        root.GET("/", Home)
        root.POST("register", Register)
        users := root.Group("users")
        {
            users.GET("/", Users)
        }
    }

    return routers;
}