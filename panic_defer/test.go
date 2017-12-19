package main

import (
	"fmt"
)

func f() {
	defer func() {
		fmt.Println("f ======= 1")
	}()

	fmt.Println("f func is running...1")
	panic(3)

	defer func() {
		fmt.Println("f ======= 2")
	}()

	fmt.Println("f func is running...2")
	fmt.Println("f func is running...3")
}

func main() {
	defer func() {
		fmt.Println("main == 1")
	}()

	defer func() {
		fmt.Println("recover == 2.1")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("recover == 2.2")
	}()

	defer func() {
		fmt.Println("main == 3")
	}()

	f()

	//发生panic之后，这个defer 就压根没入栈，也就不会执行了
	//（defer关键字的出现 以及执行 遵守 先进后出的顺序）
	defer func() {
		fmt.Println("main == 4")
	}()

	//发生panic之后，主函数中也不会再继续执行了，所以这个函数也不会被执行
	fmt.Println("end...")

}
