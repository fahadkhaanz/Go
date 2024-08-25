package cron

import (
	"fmt"
	"sync"
	"time"
)

type Cron struct {
	Updater     int
	FeederBatch int
}

func New(updater, feederBatch int) *Cron {
	return &Cron{Updater: updater, FeederBatch: feederBatch}
}

func (c *Cron) Start() {
	t := time.Now().UnixMilli()
	dataChan := make(chan int, c.FeederBatch)
	wg := sync.WaitGroup{}
	for i := 0; i < c.Updater; i++ {
		wg.Add(1)
		go c.updater(dataChan, &wg)
	}
	c.feeder(dataChan)
	close(dataChan)
	wg.Wait()
	fmt.Println("Time taken", time.Now().UnixMilli()-t, "ms")
}

func (c *Cron) feeder(dataChan chan int) {
	fmt.Println("Feeding data")
	for i := 0; i < 1000000; i++ {
		dataChan <- i
	}
}

func (c *Cron) updater(dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range dataChan {
		fmt.Println("Updating data", val)
	}
}
