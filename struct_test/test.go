package main

import ()
import "fmt"

type xSkills []string
type xHuman struct {
	name   string
	age    int
	weight int
}
type xStudent struct {
	xHuman     // 匿名字段，struct
	xSkills    // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

type skills []string

type structTest struct {
	name string
	sex  bool
	age  int
}

//在结构体中，匿名类型可以被继承，test02Struct 可以直接访问 structTest中的 name，sex，age
type test02Struct struct {
	structTest
	skills
	age int //如果 字段名 和匿名类型中的 字段 冲突了，则go实现了 外层对内层的覆盖，实现了对基类的覆写
	//但还是可以通过基类的名字去访问 内层该字段的值
}

//自定义类型中添加方法，添加函数
func (a test02Struct) Less(b test02Struct) bool {
	return a.age < b.age
}

func main() {
	var sdx structTest
	sdx1 := []string{"sdf"}
	sdx1 = append(sdx1, "xxx", "ttt")
	sda := test02Struct{sdx, sdx1, 1}
	sdb := test02Struct{sdx, sdx1, 2}

	fmt.Printf("%#v\n", sda)
	fmt.Println(sda.Less(sdb))

}
