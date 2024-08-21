package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title"`
    Description string             `bson:"description"`
    Director    string             `bson:"director"`
    ReleaseYear int                `bson:"release_year"`
}
