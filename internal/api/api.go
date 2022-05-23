package api

import (
	"github.com/duxianghua/gin-standard-template/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	router.Register(r)
	r.Run()
}
