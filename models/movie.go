package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    Director    string             `bson:"director" json:"director"`
    ReleaseYear int                `bson:"release_year" json:"release_year"`
}
