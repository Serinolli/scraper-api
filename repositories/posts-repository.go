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
