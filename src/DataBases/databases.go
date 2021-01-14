package DataBases

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "github.com/spf13/viper"
)

var Db *gorm.DB


func init()  {
    var err error

    // read mysql information from config file
    viper.SetConfigName("config")
    viper.AddConfigPath("config")
    viper.SetConfigType("json")
    err = viper.ReadInConfig()
    if err != nil {
        fmt.Printf("config file error: %s\n", err)
    }

    // concatenate string to build params of mysql connection
    username := viper.GetString(`mysql.username`)
    url := viper.GetString(`mysql.url`)
    password := viper.GetString(`mysql.password`)
    database := viper.GetString(`mysql.database`)
    mysql_params := username + ":" + password + "@tcp" + "(" + url + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

    // connect mysql
    Db, err = gorm.Open("mysql", mysql_params)
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if Db.Error != nil {
        fmt.Printf("database error %v", Db.Error)
	}
}