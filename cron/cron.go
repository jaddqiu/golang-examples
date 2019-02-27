// example use of cron package
package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

// main function for cron
func main() {
	fmt.Println("crontab schedule starts")
	c := cron.New()
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() { fmt.Println("Every 1 seconds") })
	spec = "*/10 * * * * *"
	c.AddFunc(spec, func() { fmt.Println("Every 10 seconds") })
	c.Start()
	time.Sleep(time.Second * 60)
	c.Stop()
	fmt.Println("crontab schedule ends")
}
