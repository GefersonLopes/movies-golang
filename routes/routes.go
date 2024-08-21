package routes

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateRoutes(router *mux.Router, client *mongo.Client) {
	MovieRouters(router, client)
}