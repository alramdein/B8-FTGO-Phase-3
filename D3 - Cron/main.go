package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// running setiap jam 20:40 setiap hari
	c.AddFunc("40 20 * * *", func() {
		fmt.Println("Running job at", time.Now())
	})

	// kalo mau setiap 1 menit jadi gini: */1 * * * *

	c.Start()

	// run the program indefinitely
	select {}
}
