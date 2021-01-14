package routers

import(
	"github.com/gin-gonic/gin"
    ."github.com/KuangjuX/Lang-Huan-Blessed-Land/Controllers"
)

func InnitRouter() *gin.Engine{
    routers := gin.Default()
    root := routers.Group("/")
    {   
        root.GET("/", Home)
        root.POST("register", Register)
        root.POST("loginByUsername", LoginByUsername)
        root.POST("loginByEmail", LoginByEmail)
        // users := root.Group("users")
        // {
            
        // }
    }

    return routers;
}