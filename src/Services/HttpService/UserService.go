package HttpService

import(
	"fmt"
	"errors"
	"net/http"
	"net/url"
	"io/ioutil"

	"github.com/spf13/viper"
)

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


func RequestGithubAuth()(error){
	client_id, _, _ := GetGithubSecret()
	if len(client_id) == 0{
		return errors.New("Not Found Client ID")
	}

	redirect_url := "http://langhuan.kuangjux.top/api/OAuth/github/redirect"
	url := "http://langhuan.kuangjux.top/api/OAuth/github?" + "client_id=" + client_id + "redirect_uri" + redirect_url
	response, err := http.Get(url)
	if err != nil{
		return err
	}
	defer response.Body.Close()
	return nil
}

func RequestGithubToken(code string)(string, error){
	client_id, client_secret, _ := GetGithubSecret()
	if len(client_id) == 0 || len(client_secret) == 0{
		return "", errors.New("Not Found Client ID or Secret.")
	}

	response, err := http.PostForm("https://github.com/login/oauth/access_token",
					url.Values{
						"client_id": {client_id},
						"client_secret": {client_secret},
						"code": {code},
					})
	if err != nil{
		return "", err
	}


	res, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return "", err
	}
	defer response.Body.Close()
	return string(res), nil
}
