/*
如何理解这种现象呢，可以这么理解：
所有的goroutine，他们都直接归属于main管理，生成了就挂靠在main名义下，没有什么儿子的儿子goroutine
都是main 的儿子，只要main没有结束，还在等待，goroutine就会持续运行下去
*/

package main

import ()

//import "sync"
import "fmt"
import "time"

//var wg sync.WaitGroup
//var wg2 sync.WaitGroup

func loop1() {
	//defer wg.Done()
	index := 0
	for {
		fmt.Println("s-d-f---1---", index)
		time.Sleep(1 * time.Second)
		index++
		if index > 10 {
			break
		}
	}
	fmt.Println("loop1结束")
}

func loop2() {
	//defer wg.Done()
	index := 0
	for {
		fmt.Println("s-d-f---2---", index)
		time.Sleep(1 * time.Second)
		index++
		if index > 8 {
			break
		}
	}
	fmt.Println("loop2结束")
}

func sdf() {
	//wg.Add(2)
	go loop1()
	go loop2()
	//wg.Wait()
	fmt.Println("sdf finished.0sd0f")
}

func main() {
	go sdf()
	time.Sleep(1000 * time.Second)
}
