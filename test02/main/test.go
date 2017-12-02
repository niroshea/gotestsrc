package main

// import "fmt"

// //Add xx
// func Add(x, y int, ch chan int) {
// 	ch <- 1
// 	z := x + y
// 	fmt.Println("result: ", z)
// }

// func main() {
// 	chs := make([]chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		chs[i] = make(chan int)
// 		fmt.Println("------------")
// 		go Add(i, i, chs[i])
// 	}

// 	for _, ch := range chs {
// 		<-ch
// 	}

// }
