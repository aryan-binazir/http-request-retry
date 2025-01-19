package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aryan-binazir/http-request-retry/v2/internal/database/mongodb"
	"github.com/aryan-binazir/http-request-retry/v2/internal/retryMechanism"
)

var wg sync.WaitGroup

func main() {
	startService := func(name string, fn func() error, wg *sync.WaitGroup) {
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

	mongodb.CreateConnection()
	go startService("Retry Service", retrymechanism.Init, &wg)
	select {}
}
