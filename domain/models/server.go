package model

import "go.mongodb.org/mongo-driver/mongo"

type Server struct {
	Client *mongo.Client
}
