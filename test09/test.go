package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	square := make(chan int)

	go func() {
		for i := 1; i < 4; i++ {
			naturals <- i
		}
		defer close(naturals)
	}()

	go func() {
		for x := range naturals {
			square <- x * x
		}
		close(square)
	}()

	for x := range square {
		fmt.Printf("the results :%d\n", x)
		time.Sleep(1 * time.Second)
	}
}
