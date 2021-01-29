package config

import(
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
	DataBase int
}

type GithubOAuth struct{
	ClientID string
	ClientSecret string
}

var Mysql Mysql
var Redis Redis
var GithubOAuth GithubOAuth

func init() {
	
}