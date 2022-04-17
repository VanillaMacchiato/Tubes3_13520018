package v1

import "github.com/gin-gonic/gin"

func GetResult(router *gin.RouterGroup) {
	router.GET("/result")	
}