package v1

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func PredictPatientRoute(router *gin.RouterGroup) {
	router.POST("/predict-patience", controllers.PredictPatientController())
}