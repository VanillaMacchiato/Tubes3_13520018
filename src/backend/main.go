package main

import (
	"backend/db"
	v1 "backend/routes/v1"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
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
	router.GET("/", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", os.Getenv("FE_URL"))
		ctx.JSON(http.StatusOK, "DNA at Work! Backend")
	})

	v1Group := router.Group("/api/v1")
	{
		v1.AddDiseaseRoute(v1Group)
		v1.PredictPatientRoute(v1Group)
		v1.GetResultRoute(v1Group)
	}

	router.Use(cors.Default())
	router.Run()
}
