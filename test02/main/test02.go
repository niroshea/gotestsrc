package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var values []string

func printPrime(str string) {
	defer wg.Done()

NEXT:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue NEXT
			}
		}
		fmt.Printf("%s:%d\n", str, outer)
		value := strconv.Itoa(outer)
		values = append(values, str, ":", value, "\n")
	}
	fmt.Printf("Terminate Program %s", str)
	values = append(values, "Terminate Program ", str, "\n")
}

func writeValues(values []string, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		//str := strconv.Itoa(value)
		file.WriteString(value)
	}
	return nil
}

func main() {
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println(runtime.NumCPU)
	//
	wg.Add(2)

	//values = make([]string, 0)

	fmt.Println("Start Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish.")

	wg.Wait()

	fmt.Println("\nTerminating Program ALL")
	writeValues(values, "sdf.txt")
}
