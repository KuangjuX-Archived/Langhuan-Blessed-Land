package routers

import(
	"github.com/gin-gonic/gin"
    ."github.com/KuangjuX/Lang-Huan-Blessed-Land/Controllers"
    "github.com/KuangjuX/Lang-Huan-Blessed-Land/Middleware"
)

func InnitRouter() *gin.Engine{
    routers := gin.Default()
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
        }
    }

    return routers;
}