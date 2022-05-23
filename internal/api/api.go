package api

import (
	"github.com/duxianghua/gin-standard-template/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Service() *gin.Engine {
	r := gin.Default()
	router.Register(r)
	return r
}
