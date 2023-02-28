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
	//定义两个匿名函数
	add := func(n1, n2 int) int {
		return n1 + n2
	}

	mult := func(n1, n2 int) int {
		return n1 * n2
	}
	//结果是一致的，但是不常用
	rt := calc(1, 2, add)
	fmt.Println(rt)
	rt2 := calc(3, 4, mult)
	fmt.Println(rt2)
}
