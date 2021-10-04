package Controllers

import (
	"strconv"

	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/json"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/gin-gonic/gin"
)

func GetAllArticles(c *gin.Context) {
	articles, err := Models.GetAllArticles()
	if err != nil {
		json.JsonError(c, err)
	}

	json.JsonDataWithSuccess(c, articles)
}

func GetArticlesByTag(c *gin.Context) {
	tag_id := c.Query("tag_id")
	data, err := Models.GetArticlesByTag(tag_id)

	if err != nil {
		json.JsonError(c, err)
	}

	json.JsonDataWithSuccess(c, data)
}

func GetAllArticlesByPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	params := make(map[string]string)

	if user_id, is_exist := c.GetQuery("user_id"); is_exist == true {
		params["user_id"] = user_id
	}

	if tag_id, is_exist := c.GetQuery("tag_id"); is_exist == true {
		params["tag_id"] = tag_id
	}

	data, err := Models.GetArticlesByPage(page, size, params)

	if err != nil {
		json.JsonError(c, err)
	}

	json.JsonDataWithSuccess(c, data)

}

func SearchArticles(c *gin.Context) {
	search_text := c.Query("search_text")
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	data, err := Models.SearchArticles(search_text, page, size)
	if err != nil {
		json.JsonMsgWithError(c, "Fail to search content", err)
		return
	}

	json.JsonDataWithSuccess(c, data)
}
