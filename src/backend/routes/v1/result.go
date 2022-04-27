package v1

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func GetResult(router *gin.RouterGroup) {
	router.GET("/result", controllers.GetResultController())
}
