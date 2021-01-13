package api

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Controllers"
)

func Home(c *gin.Context){
    c.String(http.StatusOK, Controllers.HomeContent())
}