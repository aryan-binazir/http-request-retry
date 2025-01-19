package mongodb

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client

func CreateConnection() error {
	var uri string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}
	fmt.Println("Connecting to MongoDB...")

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var result bson.M

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	Client = client
	return nil
}

// GetClient returns the MongoDB client instance
func GetClient() *mongo.Client {
	if Client == nil {
		if err := CreateConnection(); err != nil {
			fmt.Errorf("Error creating connection: %v", err)
		}
	}
	return Client
}
