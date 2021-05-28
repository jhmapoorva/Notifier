package main

import (
	"GoVaccineUpdaterNotifier/Service"
	"context"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	//TODO: Remove loadEnv() in production, and set env variables
	loadEnv()
	port := getEnv("PORT", "50051")
	host := "localhost"
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()
	mongoURI := options.Client().ApplyURI(
		"mongodb+srv://" + getEnv("MONGODB_USER", "NONE") + ":" + getEnv("MONGODB_PASSWORD", "NONE") + "@" + getEnv("MONGODB_URL", "NONE") + "/notified?retryWrites=true&w=majority",
	)
	mongoClient, err := mongo.NewClient(mongoURI, &options.ClientOptions{
		Registry: reg,
	})
	if err != nil {
		log.Fatalln("Unable to create new mongo client:", err)
	}
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatalln("Unable to connect to db:", err)
	}
	gRPCServer := grpc.NewServer()
	server := Service.NewServer(mongoClient)
	Service.RegisterEndpointsServer(gRPCServer, &server)
	go func() {
		if err := gRPCServer.Serve(lis); err != nil {
			log.Fatalln("Failed to serve:", err)
		}
	}()
	log.Println("Server successfully started on port:", port)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Stopping server")
	gRPCServer.Stop()
	log.Println("Stopped server successfully")
}

func getEnv(key string, defaultValue string) (retVal string) {
	retVal = os.Getenv(key)
	if len(retVal) == 0 {
		retVal = defaultValue
		log.Println(key, "not set. Using default value", defaultValue)
	}
	return
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
