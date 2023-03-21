package router

import (
	"jan-galek/api/src/endpoints"
)
import "github.com/gin-gonic/gin"

func HandleRequest() *gin.Engine {
	myRouter := gin.Default()
	myRouter.GET("/", endpoints.Homepage)

	api := myRouter.Group("/api")

	getV1("/v1", api)

	return myRouter
}
