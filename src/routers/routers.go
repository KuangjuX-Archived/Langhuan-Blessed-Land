package routers

import(
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    ."github.com/KuangjuX/Lang-Huan-Blessed-Land/Controllers"
    "github.com/KuangjuX/Lang-Huan-Blessed-Land/Middleware"
    "github.com/KuangjuX/Lang-Huan-Blessed-Land/Services/WebSocketService"
)

func InnitRouter() *gin.Engine{
    routers := gin.New()
    routers.Use(gin.Recovery())
    routers.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"x-xq5-jwt", "Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"x-xq5-jwt"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
    }))
    
    root := routers.Group("/api")
    {   
        root.GET("/", Home)
        root.POST("register", Register)
        root.POST("loginByUsername", LoginByUsername)
        root.POST("loginByEmail", LoginByEmail)
       
        
        articles := root.Group("/articles")
        {
            articles.GET("getAllArticles", GetAllArticles)
            articles.GET("getAllArticlesByPage", GetAllArticlesByPage)
            articles.GET("getArticlesByTag", GetArticlesByTag)
            articles.GET("searchArticles", SearchArticles)
            articles.GET("getCommentsByArticle",GetCommentsByArticle)
        }

        user := root.Group("user")
        {
            user.Use(Middleware.UserAuth)
            user.GET("getUserArticles", GetUserArticles)
            user.POST("modifyNickname", ModifyNickname)
            user.POST("modifyEmail", ModifyEmail)
            user.POST("modifyAvatar", ModifyAvatar)
            user.POST("uploadArticle", UploadArticle)
            user.POST("deleteArticle", DeleteArticle)
            user.POST("modifyArticle", ModifyArticle)
            user.POST("likeArticle", LikeArticle)
            
            //websocket test
            userchat := user.Group("chat")
            {
                hub := WebSocketService.NewHub()
                go hub.Run()
                userchat.GET("/ws", Chat)
                userchat.GET("/getChatInfo", GetChatInfo)
            }
        }
    }


    return routers;
}