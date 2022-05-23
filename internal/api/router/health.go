package router

import (
	"github.com/duxianghua/gin-standard-template/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func healthApi(r *gin.Engine) {
	r.GET("/ping", (&handler.HealthHandler{}).Ping)
	health := r.Group("_health") // {} must new line
	{
		health.GET("/liveness", (&handler.HealthHandler{}).Ping)
		health.GET("/readiness", (&handler.HealthHandler{}).Ping)
	}
}
