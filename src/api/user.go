package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
)

func Register(c *gin.Context)  {
    username := c.PostForm("username")
    password := c.PostForm("password")
    email    := c.PostForm("email")

    error_code := Models.CreatUser(username, password, email)

    if error_code == 1 {
        c.JSON(http.StatusOK, gin.H{
        "error_code": 1,
        "message": "创建失败",
    })
    }else{
        c.JSON(http.StatusOK, gin.H{
            "error_code": 0,
            "message": "创建成功",
    })
    }
}

func Users(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
        "code" : 1,
        "data" : "用户列表",
    })
}