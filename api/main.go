package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	r "github.com/Serinolli/scraper-api/repositories"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var port string = "8000"

func main() {
	muxRouter := mux.NewRouter()

	fmt.Println("Stablishing connection with the MongoDB database...")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	server := &r.Server{Client: client}
	fmt.Println("Listening to port " + port)
	muxRouter.HandleFunc("/posts", server.GetAllPosts).Methods("GET")
	muxRouter.HandleFunc("/posts", server.CreatePosts).Methods("POST")
	muxRouter.HandleFunc("/post", server.CreatePost).Methods("POST")
	//muxRouter.HandleFunc("/posts/{id}", getPost).Methods("GET")
	//muxRouter.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	//muxRouter.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, muxRouter))
}
