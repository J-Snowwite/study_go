package main

import "fmt"

func calc(n1, n2 int, callback func(int, int) int) int {
	//不定义是什么运算
	//通过函数参数传递我要进行的运算
	rt := callback(n1, n2)
	//检查结果在 0--100范围内，超过 -1
	if rt >= 0 && rt <= 100 {
		return rt

	}
	return -1
}

func main() {

	//直接将函数放到参数中，直接使用--更像是匿名函数
	rt := calc(1, 2, func(n1, n2 int) int {
		return n1 + n2
	})
	fmt.Println(rt)
	rt2 := calc(3, 4, func(n1, n2 int) int {
		return n1 * n2
	})
	fmt.Println(rt2)

	//变种方法调用--直接调用--主要因为函数作用域--一些值的初始化
	func() { fmt.Println(1) }()
}
