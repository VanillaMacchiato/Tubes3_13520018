package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Disease struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	DNA  string             `bson:"dna,omitempty"`
}

type Result struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Date        primitive.DateTime `bson:"date,omitempty"`
	PatientName string             `bson:"patientName,omitempty"`
	DiseaseName string             `bson:"diseaseName,omitempty"`
	HasDisease  bool               `bson:"hasDisease"`
	Likeness    float32            `bson:"likeness"`
}
