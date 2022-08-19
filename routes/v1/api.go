package v1

import "github.com/gin-gonic/gin"

func ApiRoutes(router *gin.RouterGroup) {

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "Hurry server working............")
	})
}
