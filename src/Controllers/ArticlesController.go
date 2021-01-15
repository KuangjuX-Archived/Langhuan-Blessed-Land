package Controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
)

func GetAllArticles(c *gin.Context){
	articles, err := Models.GetAllArticles()
	if err != nil {
		Help.JsonError(c, err)
	}

	Help.JsonDataWithSuccess(c, articles)
}

func GetArticleByTag(c *gin.Context){
	tag_id := c.Query("tag_id")
	data, err := Models.GetArticleByTag(tag_id)

	if err != nil {
		Help.JsonError(c, err)
	}

	Help.JsonDataWithSuccess(c, data)
}