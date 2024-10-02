package main

import (
	"fmt"
	"time"

	"main.go/ratelimiter"
)

func main() {
	rl := ratelimiter.NewRateLimiter(5, 1)
	t := time.Now().UnixMilli()
	request := 0
	totalRequest := 100000
	for i := 0; i < totalRequest; i++ {
		if rl.IsAllowed("client1") {
			fmt.Println("client1 is allowed ", i, time.Now().UnixMilli())
			request++
		}
	}
	fmt.Println("Time taken ", time.Now().UnixMilli()-t)
	fmt.Printf("Request made %d out of %d ", request, totalRequest)
}
