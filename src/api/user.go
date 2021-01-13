package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Home(c *gin.Context){
    c.String(200, "Hello, Lang Huan Blessed Land!")
}
func Users(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
        "code" : 1,
        "data" : "用户列表",
    })
}