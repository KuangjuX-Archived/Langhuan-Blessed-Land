package postgresql

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    . "github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/config"
)

var PostgreDB *gorm.DB

func init() {
	username := POSTGRE.Username
	password := POSTGRE.Password
	url := POSTGRE.URL 
	database := POSTGRE.Database

	dsn := "host=" + url + "user=" + username + "password=" + password +  "dbname=" + database +  "port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	PostgreDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if PostgreDB.Error != nil {
        fmt.Printf("database error %v", PostgreDB.Error)
	}
}