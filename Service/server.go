package Service

import (
	"GoVaccineUpdaterNotifier/DataStructures"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	MongoClient     *mongo.Client
	AlreadyNotified DataStructures.Set
}

func (s Server) mustEmbedUnimplementedEndpointsServer() {
	return
}

func NewServer(MongoClient *mongo.Client) Server {
	s := Server{}
	s.MongoClient = MongoClient
	s.AlreadyNotified = DataStructures.NewSet()
	return s
}
