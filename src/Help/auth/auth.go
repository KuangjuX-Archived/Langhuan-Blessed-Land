package auth

import(
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
)


func GetUserByToken(c *gin.Context) (interface{}, error){
	v,exist := c.Get("User")
	if !exist {
		return nil, errors.New("User not exist")
	}
	user, ok := v.(Models.User)
	if ok {
		return user, nil
	}
	return nil, errors.New("can't convert to user struct")
}