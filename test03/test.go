package main

import (
	"fmt"
	"sync"
)

func xtest(s1 string, s2 string) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(str string) {
		defer wg.Done()

	NEXT:
		for outer := 2; outer < 100; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue NEXT
				}
			}
			fmt.Printf("%s:%d\n", str, outer)
		}
		fmt.Printf("Terminate Program %s\n", str)
	}("A")

	go func(str string) {
		defer wg.Done()

	NEXT:
		for outer := 2; outer < 50; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue NEXT
				}
			}
			fmt.Printf("%s:%d\n", str, outer)
		}
		fmt.Printf("Terminate Program %s\n", str)
	}("B")

	// go func() {
	// 	//defer wg.Done()

	// 	fmt.Println("start to wait...")
	// 	wg.Wait()
	// 	fmt.Println("stop to wait...")
	// }()
	wg.Wait()
	fmt.Println("Finished!")
}

func main() {
	//var wg1 sync.WaitGroup
	//wg1.Add(0)

	//go
	xtest("A", "B")

	//wg1.Wait()
	fmt.Println("======================")
}
