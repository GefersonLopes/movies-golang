package routes

import (
	"github.com/gorilla/mux"
	"golang-movie/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func MovieRouters(router *mux.Router, client *mongo.Client) {
	router.HandleFunc("/movies", controllers.CreateMovie(client)).Methods("POST")
	router.HandleFunc("/movies", controllers.GetMovies(client)).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovie(client)).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie(client)).Methods("DELETE")
}