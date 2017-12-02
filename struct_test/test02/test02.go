package main

import (
	"fmt"
)

type kix interface {
	write(a, b string) string
	read(a, b string) string
}
type kix1 interface {
	write(a, b string) string
}

type kix2 interface {
	read(a, b string) string
}

type fix struct {
}

func (f fix) write(a, b string) string {
	fmt.Println("write a file and return 'write'")
	return "write"
}

func (f fix) read(a, b string) string {
	fmt.Println("read a file and return 'read'")
	return "read"
}

func main() {
	var a kix = new(fix) //new 中放置的是类型（自定义的或者内置的，一般为自定义的）
	var b kix1 = new(fix)
	var c kix2 = new(fix)

	a.read("xx", "xx")
	b.write("xxx", "xxx")
	c.read("xxxx", "xxxx")
}
