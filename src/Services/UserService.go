package Services

import (
	"errors"

	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/gin-gonic/gin"
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
	return nil , errors.New("can't convert to user struct")
}