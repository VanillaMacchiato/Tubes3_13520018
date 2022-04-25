package main

import (
	"backend/db"
	v1 "backend/routes/v1"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// gin router
	router := gin.Default()

	// connect to database
	db.Connect()

	// ping
	router.GET("/", func(ctx *gin.Context) {ctx.JSON(http.StatusOK, "DNA at Work! Backend")})

	v1Group := router.Group("/api/v1")
	{
		v1.AddDiseaseRoute(v1Group)
		v1.PredictPatience(v1Group)
		v1.GetResult(v1Group)
	}

	router.Run()
}
