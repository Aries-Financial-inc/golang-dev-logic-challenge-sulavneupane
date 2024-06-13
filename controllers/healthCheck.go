package controllers

import "github.com/gin-gonic/gin"

type HealthCheckService struct{}

func (healthCheckService HealthCheckService) RegisterRoutes(router *gin.Engine) {
	router.GET("/health-check", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.String(200, "OK")
}
