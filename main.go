package main

import (
	"context"
	"log"
	"net/http"

	"golang-movie/database"
    "golang-movie/routes"
	"github.com/gorilla/mux"
)

func main() {
    client := db.ConnectDB()
    defer client.Disconnect(context.Background())

    router := mux.NewRouter()
    routes.GenerateRoutes(router, client)

    log.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
