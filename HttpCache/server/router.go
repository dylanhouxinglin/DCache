package server

import "github.com/gin-gonic/gin"

func SetupRouter(request *gin.Engine) *gin.Engine {
	request.GET("/ready", func(context *gin.Context) {
		context.String(200, "success")
	})
	request.GET("/get", cacheGet)
	return request
}