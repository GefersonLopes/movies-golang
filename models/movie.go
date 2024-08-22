package models

type Movie struct {
    ID          int64              `bson:"_id,omitempty"`
    Title       string             `bson:"title"`
    Description string             `bson:"description"`
    Director    string             `bson:"director"`
    ReleaseYear int                `bson:"release_year"`
}
