package main

import "fmt"

func main() {
	isBool := true //bool类型的数据
	fmt.Printf("%T,%#v\n", isBool, isBool)

	//操作--逻辑运算
	a, b, c, d := true, true, false, false

	//逻辑与：左操作数和右操作数都为true，结果为true  &&
	println("a,b:", a && b)
	println("a,c:", a && c)
	println("c,b:", c && b)
	println("c,d:", c && d)

	//逻辑或：左操作数和右操作数只要一个为true，结果为true ||
	println("a,b:", a || b)
	println("a,c:", a || c)
	println("c,b:", c || b)
	println("c,d:", c || d)

	//逻辑非： 取反，true => false  false => true  !
	println("a:", !a)
	println("c:", !c)

	//关系操作
	fmt.Println(a == b) //打印出 true 等于 true 的值
	fmt.Println(a != b) //打印出 true 不等于 ture 的值
	fmt.Println(a == c) //打印出 true 等于 false 的值
	fmt.Println(c != b) //打印出   false 不等于 ture 的值

	//布尔类型的初始化值为 false
	var AAA bool
	fmt.Println(AAA)
}
