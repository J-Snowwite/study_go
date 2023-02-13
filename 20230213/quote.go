package main

import "fmt"

func test1(n int) {
	n = 1
	//值类型是两个不同的地址中存储的值，所以不影响
}

func test2(s []int) {
	fmt.Printf("%p\n", s)
	//打印出来后和  b   的值一致
	s[0] = 1
	//引用类型是两个不同变量存储的指向同一个地址的值，所以影响
}

func main() {

	a := 0
	b := make([]int, 10)
	test1(a)
	test2(b)
	//传递值类型参数，不会有影响
	fmt.Println(a)
	//传递引用类型参数，函数内部赋值会有影响
	fmt.Println(b)
	fmt.Printf("%p\n", b)
}
