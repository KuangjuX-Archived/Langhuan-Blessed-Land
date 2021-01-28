package Controllers

import(
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Services/HttpService"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/json"
)

func GetCommentsByArticle(c *gin.Context){
	article_id, _ := strconv.Atoi(c.Query("article_id"))
	
	data, err := HttpService.BFSComments(article_id)
	if err != nil{
		json.JsonError(c, err)
		return
	}

	json.JsonDataWithSuccess(c, data)
}