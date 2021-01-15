package Controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Services"
)

func GetArticlesByTag(c *gin.Context){
	user, err := Services.GetUserByToken(c)
	if err != nil {
		Help.JsonError(c, err)
	}

	Help.JsonDataWithSuccess(c, user)
}