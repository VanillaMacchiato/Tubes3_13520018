package controllers

import (
	"backend/db"
	"backend/lib/algorithms"
	"backend/lib/dna"
	"backend/models"
	"backend/utils/status"
	"backend/utils/status/statuscodes"
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var resultsCollection *mongo.Collection = db.GetCollection(db.DB, db.COLLECTION_RESULTS)

func PredictPatientController() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// read text file
		file, _, err := c.Request.FormFile("file")
		patientName := c.Request.FormValue("name")
		diseasePrediction := c.Request.FormValue("disease")

		if err != nil {
			// error on receiving file
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.FileError, "message": statuscodes.Text(statuscodes.FileError)})
			return
		}
		defer file.Close()

		// Check patient name
		if len(patientName) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.PatientNameEmpty, "message": statuscodes.Text(statuscodes.PatientNameEmpty)})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": statuscodes.Text(statuscodes.ServerError)})
		}

		// Check if the disease name exist
		var diseaseDetail models.Disease
		err = diseasesCollection.FindOne(ctx, bson.M{"name": diseasePrediction}).Decode(&diseaseDetail)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.DiseaseNotExits, "message": statuscodes.Text(statuscodes.Text(statuscodes.DiseaseNotExits))})
			return
		}

		// create buffer for text file
		buf := bytes.NewBuffer(nil)
		if _, err = io.Copy(buf, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
			return
		}

		// validate DNA in the content
		if !dna.ValidateDNA(buf.String()) {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.InvalidDNA, "message": statuscodes.Text(statuscodes.InvalidDNA)})
			return
		}

		// Check if the patient's DNA matches the predicted disease
		matches := algorithms.KMP(buf.String(), diseaseDetail.DNA)
		found := false
		if len(matches) > 0 {
			found = true
		}

		// Add new result
		newResult := models.Result{
			ID:          primitive.NewObjectID(),
			Date:        primitive.NewDateTimeFromTime(time.Now()),
			PatientName: patientName,
			DiseaseName: diseasePrediction,
			HasDisease:  found,
		}
		
		// Add result to DB
		_, err = resultsCollection.InsertOne(ctx, newResult)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": status.Success,
			"message": statuscodes.PredictSuccess,
			"data": gin.H{
				"_id":         newResult.ID,
				"date":        newResult.Date,
				"patientName": newResult.PatientName,
				"diseaseName": newResult.DiseaseName,
				"hasDisease":  newResult.HasDisease,
		}})

	}

}
