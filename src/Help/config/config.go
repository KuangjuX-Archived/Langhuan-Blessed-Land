package config

import(
	"fmt"

	"github.com/spf13/viper"
)

type Mysql struct{
	URL string 
	Username string
	Password string
	Database string
}

type Redis struct{
	Server string
	Password string
	Database int
}

type GithubOAuth struct{
	ClientID string
	ClientSecret string
}

var MysqlConfig Mysql
var RedisConfig Redis
var GithubOAuthConfig GithubOAuth

func init() {
	// read mysql information from config file
    viper.SetConfigName("config")
    viper.AddConfigPath("../config")
    viper.SetConfigType("json")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("config file error: %s\n", err)
	}
	
	MysqlConfig.URL = viper.GetString(`mysql.url`)
	MysqlConfig.Username = viper.GetString(`mysql.username`)
	MysqlConfig.Password = viper.GetString(`mysql.password`)
	MysqlConfig.Database = viper.GetString(`mysql.database`)

	RedisConfig.Server = viper.GetString(`redis.server`)
	RedisConfig.Password = viper.GetString(`redis.password`)
	RedisConfig.Database = viper.GetInt(`redis.database`)

	GithubOAuthConfig.ClientID = viper.GetString(`github-oauth.client_id`)
	GithubOAuthConfig.ClientSecret = viper.GetString(`github-oauth.client_secret`)

}