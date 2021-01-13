package main

import (
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/routers"
	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

func main() {
	defer orm.Db.Close()
	r := routers.InnitRouter();
    r.Run(":8081");
}