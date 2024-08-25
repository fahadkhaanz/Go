package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"main.go/cron"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	fmt.Println("Starting Main	")
	cron := cron.New(10000, 100000)
	cron.Start()
}
