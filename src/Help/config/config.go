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

type Postgre struct {
	URL string
	Username string 
	Password string 
	Database string
}

var MYSQL Mysql
var REDIS Redis
var GITHUB GithubOAuth
var POSTGRE Postgre

func init() {
	// read mysql information from config file
    viper.SetConfigName("config")
    viper.AddConfigPath("config")
    viper.SetConfigType("json")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("config file error: %s\n", err)
	}
	
	// Get Config Information
	MYSQL.URL = viper.GetString(`mysql.url`)
	MYSQL.Username = viper.GetString(`mysql.username`)
	MYSQL.Password = viper.GetString(`mysql.password`)
	MYSQL.Database = viper.GetString(`mysql.database`)

	POSTGRE.URL = viper.GetString(`postgresql.url`)
	POSTGRE.Username = viper.GetString(`postgresql.username`)
	POSTGRE.Password = viper.GetString(`postgresql.password`)
	POSTGRE.Database = viper.GetString(`postgresql.database`)

	REDIS.Server = viper.GetString(`redis.server`)
	REDIS.Password = viper.GetString(`redis.password`)
	REDIS.Database = viper.GetInt(`redis.database`)

	GITHUB.ClientID = viper.GetString(`github-oauth.client_id`)
	GITHUB.ClientSecret = viper.GetString(`github-oauth.client_secret`)

}