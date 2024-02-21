package internal

import (
	"log"
	"runtime"

	"github.com/robfig/cron/v3"
)

func RunJob() {
	c := cron.New()
	// c.AddFunc("@every 1s", func() { log.Println("every 1s") })
	// c.AddFunc("* * * * *", func() { log.Println("every 1m") })
	c.AddFunc("2 10 * * *", func() { log.Println("10 AM") })
	c.Start()
	runtime.Goexit()
}

/*
	c := cron.New()
	c.AddFunc("30 5 * * *", RunJob)
	// c.AddFunc("@every 120s", RunJob)
	c.Start()
	runtime.Goexit()
*/
