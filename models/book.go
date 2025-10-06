package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title" validate:"required,min=2"`
	Author    string             `json:"author" bson:"author" validate:"required"`
	Publisher string             `json:"publisher" bson:"publisher" validate:"required"`
	Year      int                `json:"year" bson:"year" validate:"required,gte=1900,lte=2100"`
}
