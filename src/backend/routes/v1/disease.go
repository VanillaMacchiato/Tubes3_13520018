package v1

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AddDiseaseRoute(router *gin.RouterGroup) {
	router.POST("/add-disease", controllers.AddDiseaseController())
}
