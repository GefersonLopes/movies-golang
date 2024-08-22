package services

import (
	"context"
	"time"

	"golang-movie/middlewares"
	"golang-movie/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMovie(client *mongo.Client, movie models.Movie) (*mongo.InsertOneResult, error) {
    if err := middlewares.ValidateCreateMovie(movie); err != nil {
        return nil, err
    }
    
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

    if len(movies) == 0 {
        return []models.Movie{}, nil
    }

    return movies, nil
}

func GetMovie(client *mongo.Client, id primitive.ObjectID) (models.Movie, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    validateId := middlewares.ValidadeParamSearchMovie(id)
    if validateId != nil {
        return models.Movie{}, validateId
    }

    var movie models.Movie
    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)

    if movie.ID == primitive.NilObjectID {
        return models.Movie{}, middlewares.ReturnNotFoundMovie()
    }

    return movie, err
}

func UpdateMovie(client *mongo.Client, id primitive.ObjectID, update bson.M) (*models.Movie, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    validateId := middlewares.ValidadeParamSearchMovie(id)
    if validateId != nil {
        return nil, validateId
    }

    result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})

    if result.MatchedCount == 0 {
        return nil, middlewares.ReturnNotFoundMovie()
    }

    movie := models.Movie{}
    collection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
    
    return &movie, err

}

func DeleteMovie(client *mongo.Client, id primitive.ObjectID) (*mongo.DeleteResult, error) {
    collection := client.Database("movieDB").Collection("movies")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    validateId := middlewares.ValidadeParamSearchMovie(id)
    if validateId != nil {
        return nil, validateId
    }

    result, err := collection.DeleteOne(ctx, bson.M{"_id": id})

    if result.DeletedCount == 0 {
        return nil, middlewares.ReturnNotFoundMovie()
    }

    return result, err
}
