package main

import "main.go/cron"

func main() {
	cron := cron.New(10000, 100000)
	cron.Start()
}
