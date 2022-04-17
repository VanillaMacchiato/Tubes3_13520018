package controllers

import (
	"backend/db"
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

var diseasesCollection *mongo.Collection = db.GetCollection(db.DB, db.COLLECTION_DISEASES)

func AddDiseaseController() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// read text file
		file, _, err := c.Request.FormFile("file")
		diseaseName := c.Request.FormValue("disease-name")

		if err != nil {
			// error on receiving file
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.FileError, "message": statuscodes.Text(statuscodes.FileError)})
			return
		}
		defer file.Close()

		// Check disease name
		if len(diseaseName) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.DiseaseEmpty, "message": statuscodes.Text(statuscodes.DiseaseEmpty)})
			return
		}

		var result models.Disease
		err = diseasesCollection.FindOne(ctx, bson.M{"name": diseaseName}).Decode(&result)
		// Disease name already exists
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": status.Error, "code": statuscodes.DiseaseExists, "message": statuscodes.Text(statuscodes.DiseaseExists)})
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

		// Add disease to database
		newDisease := models.Disease{
			ID:   primitive.NewObjectID(),
			Name: diseaseName,
			DNA:  buf.String(),
		}

		_, err = diseasesCollection.InsertOne(ctx, newDisease)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": status.Error, "code": statuscodes.ServerError, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": status.Success, "code": statuscodes.DNASuccess, "message": buf.String()})
	}
}
