package main

import (
	"fmt"
)

func f() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("A")
	panic("error happend.")
	fmt.Println("B")
}

func main() {
	fmt.Println("main 1")
	defer func() {
		fmt.Println("main defer 1")
	}()

	f()

	defer func() {
		fmt.Println("main defer 2")
	}()
	fmt.Println("main 2")
}
