package main

import "fmt"

const (
	name1 string = "name1"
	name2        = "name2"
	//同样支持多个参数定义
)

func main() {
	const name string = "xjt"
	//常量定义完成后，不可以修改
	//常量定义完成后，不适应不会报错
	//常量必须指定初始化值
	//常量的其他规则和变量一致
	fmt.Println(name)

	const (
		A = "test A"
		B //常量如果第二个不赋值，那么会取上一个常量的值 (A -> B)
		C = "test C"
		D //常量如果第二个不赋值，那么会取上一个常量的值 (D -> C)
		E //常量如果第二个不赋值，那么会取上一个常量的值 (E -> D -> C)
	)
	fmt.Println(B, D, E)
}
