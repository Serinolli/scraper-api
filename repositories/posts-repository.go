package repositories

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	m "github.com/Serinolli/scraper-api/models"
	"github.com/gorilla/mux"
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

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	coll := s.Client.Database("redditscrapper").Collection("posts")
	id := (mux.Vars(r))["postId"]

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.M{"postid": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	coll := s.Client.Database("redditscrapper").Collection("posts")
	id := (mux.Vars(r))["postId"]

	var result bson.M
	err := coll.FindOneAndDelete(context.TODO(), bson.M{"postid": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	coll := s.Client.Database("redditscrapper").Collection("posts")
	id := (mux.Vars(r))["postId"]

	var post m.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result bson.M
	err := coll.FindOneAndUpdate(context.TODO(), bson.M{"postid": id}, bson.M{"$set": post}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
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

func (s *Server) CreatePosts(w http.ResponseWriter, r *http.Request) {
	coll := s.Client.Database("redditscrapper").Collection("posts")

	var posts []m.Post
	var err error

	if err := json.NewDecoder(r.Body).Decode(&posts); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var iPosts []interface{}
	for _, p := range posts {
		iPosts = append(iPosts, p)
	}

	_, err = coll.InsertMany(context.TODO(), iPosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
