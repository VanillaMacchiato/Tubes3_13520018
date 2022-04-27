package controllers

import (
	"backend/db"
	"backend/lib/searchresult"
	"backend/models"
	"backend/utils/status"
	"backend/utils/status/statuscodes"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var resultsCollection *mongo.Collection = db.GetCollection(db.DB, db.COLLECTION_RESULTS)

func GetResultController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", os.Getenv("FE_URL"))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		params := c.Request.URL.Query()

		var date, diseaseName string
		dateParam, diseaseParam, err := searchresult.SanitizeInput(params["input"][0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.WrongSearchQuery, "message": err.Error()})
			return
		}
		if dateParam != "" {
			date = dateParam
		}
		if diseaseParam != "" {
			diseaseName = diseaseParam
		}

		var cursor *mongo.Cursor
		var parsedDate time.Time

		_, offset := time.Now().Zone()
		layout := "2006-01-02"
		if date != "" {
			parsedDate, err = time.Parse(layout, date)
			parsedDate = parsedDate.Add(-time.Duration(offset) * time.Second) // time zone correction
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
				return
			}
		}
		fmt.Println(time.Duration(offset) * time.Second)

		if diseaseName != "" && date != "" {
			cursor, err = resultsCollection.Find(ctx, gin.H{"diseaseName": gin.H{"$eq": diseaseName}, "date": gin.H{"$gte": parsedDate, "$lt": parsedDate.Local().Add(time.Hour * 24)}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
				return
			}
		} else if diseaseName == "" && date != "" {
			cursor, err = resultsCollection.Find(ctx, gin.H{"date": gin.H{"$gte": parsedDate, "$lt": parsedDate.Add(time.Hour * 24)}})
		} else if date == "" && diseaseName != "" {
			fmt.Println("elif date")
			cursor, err = resultsCollection.Find(ctx, gin.H{"diseaseName": gin.H{"$eq": diseaseName}})
		} else {
			fmt.Println("else")
			cursor, err = resultsCollection.Find(ctx, gin.H{})
		}

		var res []models.Result
		if err = cursor.All(ctx, &res); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"status": status.Success, "code": statuscodes.ResultSuccess, "data": res})
		return
	}
}
