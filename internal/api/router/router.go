package router

import (
	"github.com/duxianghua/gin-standard-template/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// Unified exception handling
	r.Use(handler.Recover)
	// health apis
	healthApi(r)
	// version 1，url scale
	// http://Host:Port/api/v1/...
	api := r.Group("api")
	{
		v1 := api.Group("v1")
		{
			route_v1(v1)
		}
		v2 := api.Group("v2")
		{
			route_v2(v2)
		}
	}
}

func route_v1(rg *gin.RouterGroup) {

}

func route_v2(rg *gin.RouterGroup) {

}
