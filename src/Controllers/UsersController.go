package Controllers

import (
	"errors"

    "github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
)

func Register(c *gin.Context)  {
    username := c.PostForm("username")
    password := c.PostForm("password")
	email    := c.PostForm("email")
	
	if len(username) == 0 || len(password) == 0 || len(email) == 0 {
		Help.JsonMsgWithError(c, "Fail to register.", errors.New("Expected arguments."))
		return
	}

    message, err := Models.CreatUser(username, password, email)

    if err == nil {
       	Help.JsonMsgWithSuccess(c, message)
    }else{
    	Help.JsonMsgWithError(c, message, err)
    }
}


func LoginByUsername(c *gin.Context){
	var user Models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	data, err := user.Login()
	if err == nil{
		Help.JsonDataWithSuccess(c, data)
	}else{
		Help.JsonError(c, err)
	}
}

func LoginByEmail(c *gin.Context)  {
	var user Models.User
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	data, err := user.Login()
	if err == nil{
		Help.JsonDataWithSuccess(c, data)
	}else{
		Help.JsonError(c, err)
	}
}