package retrymechanism

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
)

func Init() error {

	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", func() { fmt.Println("Every 5 seconds") })
	c.Start()
	// c.Run()
	select {}

	return errors.New("asd")
}
