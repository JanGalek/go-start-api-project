package router

import (
	"github.com/gin-gonic/gin"
	"jan-galek/api/src/endpoints"
)

func getV1(relativePath string, api *gin.RouterGroup) *gin.RouterGroup {
	v1 := api.Group("/v1")
	v1.GET("/", endpoints.HealthCheck)
	v1.GET("/healthcheck", endpoints.HealthCheck)

	return v1
}
