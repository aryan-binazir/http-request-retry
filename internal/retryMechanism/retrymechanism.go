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
	c.AddFunc("*/5 * * * * *", writeToDB)
	c.Start()
	select {}
}

func writeToDB() {
	_, err := mongodb.NewMongoOperations().InsertOne(context.TODO(), "test", bson.M{"timestamp": time.Now(), "message": "executed"})
	if err != nil {
		log.Printf("Error inserting document: %v", err)
		return
	}

	fmt.Println("Document inserted successfully")
}
