package Controllers

import(
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Services"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
)

func GetCommentsByArticle(c *gin.Context){
	article_id, _ := strconv.Atoi(c.Query("article_id"))
	
	data, err := Services.BFSComments(article_id)
	if err != nil{
		Help.JsonError(c, err)
		return
	}

	Help.JsonDataWithSuccess(c, data)
}