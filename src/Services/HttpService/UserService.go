package HttpService

import(
	"time"
	"fmt"
	"strconv"
	"errors"
	"net/http"
	"net/url"
	"io/ioutil"

	"github.com/spf13/viper"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
)

var DuplicatedUsername error = errors.New("Username has existed. Please modify username.")
var PasswordError error = errors.New("You have changed password. Please login by current password")

func GetGithubSecret()(string, string ,error){
	viper.SetConfigName("config")
    viper.AddConfigPath("config")
	viper.SetConfigType("json")
	
	err := viper.ReadInConfig()
    if err != nil {
		fmt.Printf("config file error: %s\n", err)
		return "", "" ,err
	}
	
	client_id := viper.GetString(`github-oauth.client_id`)
	client_secret := viper.GetString(`github-oauth.client_secret`)
	return client_id, client_secret, nil
}


func RequestGithubToken(code string)([]byte, error){
	client_id, client_secret, _ := GetGithubSecret()
	if len(client_id) == 0 || len(client_secret) == 0{
		return []byte(""), errors.New("Not Found Client ID or Secret.")
	}
	
	uri := "https://github.com/login/oauth/access_token"
	data := url.Values{
				"client_id": {client_id},
				"client_secret": {client_secret},
				"code": {code},
			}

	response, err := http.PostForm(uri, data)
	response.Header.Set("Accept", "application/json")
	
	if err != nil{
		return []byte(""), err
	}


	res, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return []byte(""), err
	}
	// fmt.Printf("res: %s\n", res)
	defer response.Body.Close()
	return res, nil
}

func RequestGithubInfo(access_token string)([]byte, error){
	auth := "token " + access_token
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
    response, err := client.Do(req)
   
	if err != nil{
		return nil, err
	}

	res, err := ioutil.ReadAll(response.Body)

	if err != nil{
		return nil, err
	}

	defer response.Body.Close()

	return res, nil
}

// {"data":{"avatar_url":"https://avatars.githubusercontent.com/u/56680481?v=4","bio":"What I cannot create, I do not understand. --Richard Philip Feynman","blog":"http://mainsite.kuangjux.top/","company":"TJU","created_at":"2019-10-17T11:15:11Z","email":"qcx@tju.edu.cn","events_url":"https://api.github.com/users/KuangjuX/events{/privacy}","followers":"","followers_url":"https://api.github.com/users/KuangjuX/followers","following":"","following_url":"https://api.github.com/users/KuangjuX/following{/other_user}","gists_url":"https://api.github.com/users/KuangjuX/gists{/gist_id}","gravatar_id":"","hireable":"","html_url":"https://github.com/KuangjuX","id":"","location":"Tianjin China","login":"KuangjuX","name":"ChengXiang Qi","node_id":"MDQ6VXNlcjU2NjgwNDgx","organizations_url":"https://api.github.com/users/KuangjuX/orgs","public_gists":"","public_repos":"","received_events_url":"https://api.github.com/users/KuangjuX/received_events","repos_url":"https://api.github.com/users/KuangjuX/repos","site_admin":"","starred_url":"https://api.github.com/users/KuangjuX/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/KuangjuX/subscriptions","twitter_username":"","type":"User","updated_at":"2021-01-30T15:53:17Z","url":"https://api.github.com/users/KuangjuX"},"error_code":0}


func LoginByGithub(user_info map[string]string)(string, error){

	// Verify email is exist?
	is_email, _ := Models.IsExistUser(user_info["email"])
	
	// Veridfy username is exist?
	is_username, _ := Models.IsExistUser(user_info["login"])

	var user *Models.User
	default_password, _ := Models.Hash(user_info["login"] + strconv.FormatInt(time.Now().Unix(), 10))
	user = &Models.User{
		Username: user_info["login"],
		Nickname: user_info["name"],
		Password: string(default_password),
		Email: user_info["email"],
		Avatar: user_info["avatar_url"],
	}

	if !is_email{
		// User is not found
		if !is_username{
			fmt.Printf("123\n")
			// No duplicate username
			// Create User !
			if err := user.CreateUser(); err != nil{
				return "", err
			}
			
			token, err := user.Login()

			if err != nil{
				return "", err
			}

			return token, nil
		}else{
			// Duplicate username
			return "", DuplicatedUsername 
		}
	}else{
		// User has existed
		// Login By Default password
		token, err := user.Login()
		if err != nil{
			return "", err
		}

		return token, nil
	}

}
