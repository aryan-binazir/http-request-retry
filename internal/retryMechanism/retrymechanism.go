package retrymechanism

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aryan-binazir/http-request-retry/v2/internal/database/mongodb"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Init() error {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", writeToDB) //func() { fmt.Println("Every 5 seconds") })
	c.Start()
	select {}
}

func writeToDB() {
	// client := mongodb.GetDb()
	// collection := client.Database("test").Collection("test")

	// _, err := collection.InsertOne(context.TODO(), bson.M{"timestamp": time.Now(), "message": "executed"})
	_, err := mongodb.NewMongoOperations().InsertOne(context.TODO(), "test", bson.M{"timestamp": time.Now(), "message": "executed"})
	if err != nil {
		log.Printf("Error inserting document: %v", err)
		return
	}

	fmt.Println("Document inserted successfully")
}
