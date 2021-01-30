package HttpService

import(
	"fmt"
	"errors"
	"net/http"
	"net/url"
	"io/ioutil"

	"github.com/spf13/viper"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/debug"
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
	// var response *http.Request
	response, err := http.Get("https://api.github.com/user")
	auth := "Bearer " + access_token
	response.Header.Add("Authorization", auth)
	response.Header.Add("Accept", "application/json")
	
	debug.StdOutDebug("response: %s", response)
	debug.StdOutDebug("response header: %s", response.Header)
	debug.StdOutDebug("response body: %s", response.Body)

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
