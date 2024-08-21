package services

import (
    "context"
    "time"
    
    "golang-movie/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func CreateMovie(client *mongo.Client, movie models.Movie) (*mongo.InsertOneResult, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.InsertOne(ctx, movie)
    return result, err
}

func GetMovies(client *mongo.Client) ([]models.Movie, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var movies []models.Movie
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var movie models.Movie
        cursor.Decode(&movie)
        movies = append(movies, movie)
    }

    return movies, nil
}

func UpdateMovie(client *mongo.Client, id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
    return result, err
}

func DeleteMovie(client *mongo.Client, id primitive.ObjectID) (*mongo.DeleteResult, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
    return result, err
}
