package Controllers

import (
	"errors"
	"strconv"
	"time"

    "github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Services"
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

func GetUserArticles(c *gin.Context){
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	user, _ := Services.GetUserByToken(c)
	user_id := user.(Models.User).ID

	params := make(map[string]string)
	params["user_id"] = strconv.FormatInt(user_id, 10)

	data, err := Models.GetArticlesByPage(page, size, params)
	if err != nil {
		Help.JsonError(c, err)
		return
	}

	Help.JsonDataWithSuccess(c, data)

}

func ModifyNickname(c *gin.Context){
	new_nickname := c.PostForm("new_nickname")
	user, _ := Services.GetUserByToken(c)
	user_id := user.(Models.User).ID

	err := Models.ModifyNickname(user_id, new_nickname)
	if err != nil{
		Help.JsonError(c, err)
		return
	}
	Help.JsonSuccess(c)

}

func ModifyEmail(c *gin.Context){
	new_email := c.PostForm("new_email")
	user, _ := Services.GetUserByToken(c)
	user_id := user.(Models.User).ID

	err := Models.ModifyEmail(user_id, new_email)
	if err != nil{
		Help.JsonError(c, err)
		return
	}
	Help.JsonSuccess(c)
}

func ModifyAvatar(c *gin.Context){
	avatar, _ := c.FormFile("new_avatar")
	avatar_name := (time.Now()).String() + avatar.Filename
	avatar_path := "storage/app/image" + avatar_name

	if err := c.SaveUploadedFile(avatar, avatar_path); err != nil {
		Help.JsonMsgWithError(c, "Fail to upload image", err)
		return
	}
	
	user, _ := Services.GetUserByToken(c)
	user_id := user.(Models.User).ID
	if err := Models.ModifyAvatar(user_id, avatar_name); err != nil {
		Help.JsonMsgWithError(c, "Fail to store image into database", err)
		return
	}

	Help.JsonMsgWithSuccess(c, "Success to upload image")

}