package Help

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func JsonDataWithSuccess(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"data": data,
	})
}

func JsonError(c *gin.Context, err error){
	c.JSON(http.StatusOK, gin.H{
		"error_code": 1,
		"error": err.Error(),
	})
}

func JsonMsgWithSuccess(c *gin.Context, msg interface{}){
	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"message": msg,
	})
}

func JsonMsgWithError(c *gin.Context, msg interface{}, err error){
	c.JSON(http.StatusOK, gin.H{
		"error_code": 1,
		"message": msg,
		"error": err.Error(),
	})
}


