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

func GetArticlesByTag(c *gin.Context){
	tag_id := c.Query("tag_id")
	data, err := Models.GetArticlesByTag(tag_id)

	if err != nil {
		Help.JsonError(c, err)
	}

	Help.JsonDataWithSuccess(c, data)
}

func GetAllArticlesByPage(c *gin.Context){
	query := Models.ArticlePages{}
	err := c.ShouldBindQuery(query)

	if err != nil {
		Help.JsonError(c, err)
		return
	}
	list, total, err := query.Search()

	if err != nil {
		Help.JsonError(c, err)
		return
	}

	Help.JsonPagination(c, list, total, &query.Pagination)


}