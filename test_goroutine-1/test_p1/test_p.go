/*
用于验证 go 中 协程的非抢占式，抢时间片。
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)    //设置处理器数量，即 P 的数量
	for i := 0; i < 2; i++ { //创建两个goroutine
		go func(index int) {
			fmt.Print(index)
			for {
				if index == 1 { //模拟计算密集程序
					fmt.Print(index)
					return
				}
				//time.Sleep(1 * time.Second)
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("end")
}
