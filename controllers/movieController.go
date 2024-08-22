package controllers

import (
	"encoding/json"
	"net/http"

	"golang-movie/middlewares"
	"golang-movie/models"
	"golang-movie/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMovie(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var movie models.Movie
        _ = json.NewDecoder(r.Body).Decode(&movie)

        result, err := services.CreateMovie(client, movie)
        if err != nil {
            middlewares.HandleErros(err, w)
            return
        }

        json.NewEncoder(w).Encode(result)
    }
}

func GetMovies(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        movies, err := services.GetMovies(client)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(movies)
    }
}

func GetMovie(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, _ := primitive.ObjectIDFromHex(params["id"])

        movie, err := services.GetMovie(client, id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(movie)
    }
}

func UpdateMovie(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, _ := primitive.ObjectIDFromHex(params["id"])

        var updateData map[string]interface{}
        _ = json.NewDecoder(r.Body).Decode(&updateData)

        result, err := services.UpdateMovie(client, id, updateData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(result)
    }
}

func DeleteMovie(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id, _ := primitive.ObjectIDFromHex(params["id"])

        result, err := services.DeleteMovie(client, id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(result)
    }
}
