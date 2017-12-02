package main

import (
	"fmt"
	"sync"
)

type sx struct {
	sdf int
	str string
}

var wg sync.WaitGroup
var wg1 sync.WaitGroup

func printPrime(str string, done chan<- *sx) {
	defer wg.Done()
	fmt.Println("start the func 1", str)
NEXT:
	for outer := 2; outer < 1000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue NEXT
			}
		}
		fmt.Printf("%s:%d\n", str, outer)
		//	value := strconv.Itoa(outer)
		//		values = append(values, str, ":", value, "\n")
	}
	fmt.Printf("Terminate Program %s\n", str)
	//var x sx
	x := sx{0, str}
	done <- &x // 发送数据到通道中
}

func goRtest(str string, done chan<- *sx) {
	defer wg1.Done()
	fmt.Println("start the func 1", str)
NEXT:
	for outer := 2; outer < 1000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue NEXT
			}
		}
		fmt.Printf("%s:%d\n", str, outer)
		//	value := strconv.Itoa(outer)
		//		values = append(values, str, ":", value, "\n")
	}
	fmt.Printf("Terminate Program %s\n", str)
	//var x sx
	x := sx{189239, str + "-sdf"}
	done <- &x // 发送数据到通道中
}

func main() {

	wg.Add(2)
	wg1.Add(2)
	done := make(chan *sx, 4)
	go printPrime("A", done)
	go printPrime("B", done)
	go goRtest("C", done)
	go goRtest("D", done)
	go func() {
		//该操作会一直阻塞所在协程，直到他所等待的全部协程全部结束
		wg.Wait()

		fmt.Println("the first wait is finished!")

		wg1.Wait()

		//关闭通道之后，下面的接收操作会检测到此行为，接收操作会立即完成，不会管此goroutine 的状态
		//主goroutine的结束后自动结束该goroutine的生命
		close(done)
	}()
	//等待协程执行完毕给我信号，我将会一直阻塞 直到 通道中被传递了数据
	//xsdf := <-done
	//y := <-done

	//只有这个是不行的，只有第一个完成的goroutine会完整执行，其他都会随着主协程的死亡而死亡
	//<-done

	//这个range是一个接收通道信息的操作，在通道关闭之前他会一直监听
	//当检测到通道关闭的时候，任何接收操作都会立即完成，并且获取一个通道对应零值
	//通道中有值的时候就会执行一次，当关闭通道该操作获取了一个零值则停止监听，结束该操作，在主goroutine中往下执行
	for xs := range done {
		fmt.Printf("==============this is the end,%d,%s\n", xs.sdf, xs.str)
	}

	fmt.Println("-sd=s==sd=f=s=df=s=df=s=df=sd=f=sdf======the real end.")
}
