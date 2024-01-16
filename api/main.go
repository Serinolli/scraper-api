package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "github.com/Serinolli/scraper-api/models"
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

	fmt.Println("Listening to port " + port)
	muxRouter.HandleFunc("/posts", getPosts).Methods("GET")
	muxRouter.HandleFunc("/posts", createPost).Methods("POST")
	//muxRouter.HandleFunc("/posts/{id}", getPost).Methods("GET")
	//muxRouter.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	//muxRouter.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, muxRouter))
}

func getPosts(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode([]m.Post{{
		ID:      "sampletest1",
		Title:   "starting api",
		Upvotes: 15,
		Content: "lorem ipsum dolor sit amet",
	}})
}

func createPost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var posts []m.Post
	err := json.NewDecoder(request.Body).Decode(&posts)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(writer).Encode(posts)
}
