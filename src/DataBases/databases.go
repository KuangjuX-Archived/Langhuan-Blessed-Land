package DataBases

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql" //加载mysql驱动
    "github.com/jinzhu/gorm"
    "github.com/spf13/viper"
)

var Db *gorm.DB


func init()  {
    var err error
    viper.SetConfigName("db")
    viper.AddConfigPath("../config")
    viper.SetConfigType("json")
    err = viper.ReadInConfig()
    if err != nil {
        fmt.Printf("config file error: %s\n", err)
    }

    Db, err = gorm.Open("mysql", "root:qgKj2017@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if Db.Error != nil {
        fmt.Printf("database error %v", Db.Error)
	}
}