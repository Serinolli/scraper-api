package repositories

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	m "github.com/Serinolli/scraper-api/models"
)

type Server m.Server

func (s *Server) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	coll := s.Client.Database("redditscrapper").Collection("posts")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	coll := s.Client.Database("redditscrapper").Collection("posts")

	var post m.Post
	var err error

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = coll.InsertOne(context.TODO(), post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
