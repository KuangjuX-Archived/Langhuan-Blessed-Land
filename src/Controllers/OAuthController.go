package Controllers

import(
	"strings"
	"errors"
	"strconv"
	jsonparse "encoding/json"
	"net/http"

    "github.com/gin-gonic/gin"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/json"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/auth"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Services/HttpService"
)

var NotFoundField error = errors.New("Not Found Field.")
func handlerUser(user *map[string]string, c *gin.Context)(error){
	var keys = [...]string{"login", "name", "email", "avatar_url"}
	for _, key := range keys {
		value := c.PostForm(key)
		if len(value) > 0 {
			(*user)[key] = value
		}else{
			return NotFoundField
		}
		
	}
	return nil
}

func OAuthGithubRename(c *gin.Context) {
	user := make(map[string]string)
	err := handlerUser(&user, c)

	if err != nil{
		json.JsonError(c, err)
	}
	token, err := HttpService.LoginByGithub(user)

	if err != nil{
		if err == HttpService.DuplicatedUsername{
			json.JsonMsgWithError(c, user, err)
			return 
		}else{
			json.JsonError(c, err)
			return
		}
	}

	var ret map[string]string
	ret = make(map[string]string)
	ret["token"] = token
	json.JsonDataWithSuccess(c, ret)
}

func OAuthGithub(c *gin.Context){
	client_id, _, _ := HttpService.GetGithubSecret()
	if len(client_id) == 0{
		json.JsonError(c, errors.New("Not Found Client ID"))
		return
	}

	redirect_url := "http://127.0.0.1:8081/api/OAuth/github/redirect"
	url := "https://github.com/login/oauth/authorize?" + "client_id=" + client_id + "&redirect_uri=" + redirect_url
	
	c.Redirect(http.StatusMovedPermanently, url)
}

func OAuthGithubRedirect(c *gin.Context){
	code, ok := c.GetQuery("code")
	if !ok{
		json.JsonError(c, errors.New("Fail to get code."))
		return
	}
	response, err := HttpService.RequestGithubToken(code)
	if err != nil{
		json.JsonError(c, err)
		return
	}
	
	const prefix_len = len("access_token=")
	access_token := strings.Split(string(response), "&")[0][prefix_len:]
	

	if len(access_token) == 0 {
		json.JsonError(c, errors.New("Not Found access_token"))
		return
	}

	data, err := HttpService.RequestGithubInfo(access_token)
	if err != nil{
		json.JsonError(c, err)
		return
	}
	var res map[string]string
	jsonparse.Unmarshal(data, &res)


	token, err := HttpService.LoginByGithub(res)
	
	if err != nil{
		if err == HttpService.DuplicatedUsername{
			json.JsonMsgWithError(c, res, err)
			return 
		}else{
			json.JsonError(c, err)
			return
		}
	}

	var ret map[string]string
	ret = make(map[string]string)
	ret["token"] = token
	json.JsonDataWithSuccess(c, ret)
}

func GetUserArticles(c *gin.Context){
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	user, _ := auth.GetUserByToken(c)
	user_id := user.(Models.User).ID

	params := make(map[string]string)
	params["user_id"] = strconv.FormatInt(user_id, 10)

	data, err := Models.GetArticlesByPage(page, size, params)
	if err != nil {
		json.JsonError(c, err)
		return
	}

	json.JsonDataWithSuccess(c, data)

}