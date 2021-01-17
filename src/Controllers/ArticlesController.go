package Controllers

import(
	"strconv"

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

func GetArticlesByTag(c *gin.Context){
	tag_id := c.Query("tag_id")
	data, err := Models.GetArticlesByTag(tag_id)

	if err != nil {
		Help.JsonError(c, err)
	}

	Help.JsonDataWithSuccess(c, data)
}

func GetAllArticlesByPage(c *gin.Context){
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	params := make(map[string]string)

	if user_id, is_exist := c.GetQuery("user_id"); is_exist == true {
		params["user_id"] = user_id
	}

	if tag_id, is_exist := c.GetQuery("tag_id"); is_exist == true {
		params[tag_id] = tag_id
	}

	data, err := Models.GetArticlesByPage(page, size, params)

	if err != nil {
		Help.JsonError(c, err)
	}

	Help.JsonDataWithSuccess(c, data)

}