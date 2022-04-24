package v1

import ("github.com/gin-gonic/gin"
	"backend/controllers")

func GetResult(router *gin.RouterGroup) {
	router.GET("/result", controllers.GetResultController())	
}