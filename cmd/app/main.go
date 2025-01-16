package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/aryan-binazir/http-request-retry/v2/internal/database/mongodb"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	client, err := mongodb.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Initial commit")
	collection := client.Database("test").Collection("test")
	collection.InsertOne(context.Background(), map[string]string{"name": "abc"})
	defer client.Disconnect(context.Background())
	wg.Wait()
}
