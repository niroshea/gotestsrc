package main

import "fmt"
import "reflect"

type S struct {
	a int
}

//Set1 这就告诉我，当给结构体添加方法的时候，最好是给指针形式的
//最好也声明指针类型的来调用函数方法
func (s S) Set1(v int) {
	s.a = v
}

func (s *S) Set2(v int) {
	s.a = v
}

func (s *S) Get() int {
	return s.a
}

func main() {
	var s1 S
	var s2 *S

	s1.Set1(100) //receiver类型为T的实例，Set1修改的只是s1的副本
	fmt.Println(s1.Get())
	s1.Set2(100) //receiver类型为*T的实例，Set2修改的是s1的引用
	fmt.Println(s1.Get())
	fmt.Println("....s1 Method Set....") //receiver类型为T只包含T的方法集，receiver类型为*T包含T和*T的方法集
	DumpMethodSet(s1)
	fmt.Println("....s2 Method Set....")
	DumpMethodSet(s2)
}

//DumpMethodSet xx
func DumpMethodSet(i interface{}) {
	//反射
	MethodSet := reflect.TypeOf(i)
	for i := 0; i < MethodSet.NumMethod(); i++ {
		fmt.Println(MethodSet.Method(i).Name)
	}
}
