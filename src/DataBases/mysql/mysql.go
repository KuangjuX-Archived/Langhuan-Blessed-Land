package mysql

import (
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    . "github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/config"
)



var Db *gorm.DB
var ErrorRecordNotFound error


func init()  {
    ErrorRecordNotFound = gorm.ErrRecordNotFound 
    var err error
    
    username := MYSQL.Username
    password := MYSQL.Password
    url := MYSQL.URL
    database := MYSQL.Database
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