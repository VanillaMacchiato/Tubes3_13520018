package controllers

import (
	"backend/lib/algorithms"
	"backend/lib/dna"
	"backend/models"
	"backend/utils/status"
	"backend/utils/status/statuscodes"
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PredictPatientController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", os.Getenv("FE_URL"))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// read text file
		file, _, err := c.Request.FormFile("file")
		patientName := c.Request.FormValue("name")
		diseasePrediction := c.Request.FormValue("disease")
		algorithm := c.Request.FormValue("algorithm")

		if err != nil {
			// error on receiving file
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.FileError, "message": statuscodes.Text(statuscodes.FileError)})
			return
		}
		defer file.Close()

		// Check algorithm used for prediction
		algorithmsAvailable := map[string]bool{
			"KMP":         true,
			"BoyerMoore":  true,
			"Levenshtein": true,
			"":            true, // algorithm not specified
		}
		if !algorithmsAvailable[algorithm] {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.InvalidAlgorithm, "message": statuscodes.Text(statuscodes.InvalidAlgorithm)})
			return
		}

		// Check patient name
		if len(patientName) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.PatientNameEmpty, "message": statuscodes.Text(statuscodes.PatientNameEmpty)})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": statuscodes.Text(statuscodes.ServerError)})
			return
		}

		// Check if the disease name exist
		var diseaseDetail models.Disease
		err = diseasesCollection.FindOne(ctx, bson.M{"name": diseasePrediction}).Decode(&diseaseDetail)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.DiseaseNotExists, "message": statuscodes.Text(statuscodes.DiseaseNotExists)})
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
		found := false
		diseaseDNA := diseaseDetail.DNA
		patientDNA := buf.String()
		var likeness float32 = 0.0
		if algorithm == "KMP" {
			found = algorithms.KMP(patientDNA, diseaseDNA)
			if found {
				likeness = 100.0
			} else {
				likeness = algorithms.Likeness(buf.String(), diseaseDetail.DNA)
				if likeness >= 80.0 {
					found = true
				}
			}
		} else if algorithm == "BoyerMoore" {
			found = algorithms.BoyerMoore(patientDNA, diseaseDNA)
			if found {
				likeness = 100.0
			} else {
				likeness = algorithms.Likeness(buf.String(), diseaseDetail.DNA)
				if likeness >= 80.0 {
					found = true
				}
			}
		}

		// Add new result
		newResult := models.Result{
			ID:          primitive.NewObjectID(),
			Date:        primitive.NewDateTimeFromTime(time.Now()),
			PatientName: patientName,
			DiseaseName: diseasePrediction,
			HasDisease:  found,
			Likeness:    likeness,
		}

		// Add result to DB
		_, err = resultsCollection.InsertOne(ctx, newResult)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": status.Success,
			"code": statuscodes.PredictSuccess,
			"data": gin.H{
				"_id":         newResult.ID,
				"date":        newResult.Date,
				"patientName": newResult.PatientName,
				"diseaseName": newResult.DiseaseName,
				"hasDisease":  newResult.HasDisease,
				"likeness":    newResult.Likeness,
			}})
	}

}
