package main

import (
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/routers"
)

func main() {
	r := routers.InnitRouter();
    r.Run(":8081");
}