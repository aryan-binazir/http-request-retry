package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aryan-binazir/http-request-retry/v2/internal/database/mongodb"
	// "github.com/aryan-binazir/http-request-retry/v2/internal/retryMechanism"
)

var wg sync.WaitGroup

func main() {
	mongoConnection := func(name string, fn func() error, wg *sync.WaitGroup) {
		for {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Printf("Starting %s...\n", name)
				fn()
				fmt.Printf("%s exited. Restarting...\n", name)
			}()
			wg.Wait()
			time.Sleep(2 * time.Second)

		}

	}
	go mongoConnection("mongodb", mongodb.CreateConnection, &wg)
	select {}
	// wg.Add(1)
	// client, err := mongodb.GetClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println("Initial commit")
	// collection := client.Database("test").Collection("test")
	// // _, err = collection.InsertOne(context.Background(), map[string]string{"name": "abc"})
	// // if err != nil {
	// // 	log.Printf("Error: %v", err)
	// // 	log.Printf("Retrying...")
	// // 	_, err = collection.InsertOne(context.Background(), map[string]string{"name": "abc"})
	// // }
	//
	// defer client.Disconnect(context.Background())
	// wg.Wait()
}
