package main

import ()
import "fmt"

func main() {
	var ss []string
	fmt.Printf("length:%v \tcapcity:%v \taddr:%p \tisnil:%v \tcontent:%v\n", len(ss), cap(ss), ss, ss == nil, ss)

	//追加
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("S%d", i))
	}

	fmt.Printf("length:%v \tcapcity:%v \taddr:%p \tisnil:%v \tcontent:%v\n", len(ss), cap(ss), ss, ss == nil, ss)

	//删除
	index := 5
	ss = append(ss[:index], ss[index+1:]...)

	fmt.Printf("length:%v \tcapcity:%v \taddr:%p \tisnil:%v \tcontent:%v\n", len(ss), cap(ss), ss, ss == nil, ss)

	//插入
	rear := append([]string{}, ss[index:]...)
	fmt.Printf("content:%v\n", rear)
	ss = append(ss[:index], "SDF")
	ss = append(ss, rear...)
	fmt.Printf("length:%v \tcapcity:%v \taddr:%p \tisnil:%v \tcontent:%v\n", len(ss), cap(ss), ss, ss == nil, ss)
}
