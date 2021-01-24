package Middleware

import(
	"strings"
	"errors"

	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
	"github.com/gin-gonic/gin"

)

const contextKeyUserObj = "User"
const bearerLength = len("Bearer ")

func authentication(c *gin.Context){
	token, ok := c.GetQuery("token")
	if !ok {
		hToken := c.GetHeader("Authorization")
		if len(hToken) < bearerLength {
			Help.JsonError(c, errors.New("header Authorization has not Bearer token"))
			return 
		}
		token = strings.TrimSpace(hToken[bearerLength:])
	}

	user, err := Models.JwtParserUser(token)
	if err != nil {
		Help.JsonError(c, err)
		return
	}

	c.Set(contextKeyUserObj, *user)
	c.Next()
}

func UserAuth(c *gin.Context){
	authentication(c)
}


