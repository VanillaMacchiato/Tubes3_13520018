package v1

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func PredictPatience(router *gin.RouterGroup) {
	router.POST("/predict-patience", controllers.PredictPatientController())
}