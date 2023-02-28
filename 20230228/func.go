package main

import "fmt"

func test(a int, b string) {
	fmt.Println(a, b)
}

func test2(c, d int) int {
	return 1
}

//方法也是一种类型

func main() {
	a := test
	//将test方法赋值给a
	fmt.Printf("%T\n\n", test)
	fmt.Printf("%T\n\n", a)
	a(1, "aa")
	//调用test方法
	test(2, "bb")
	//函数是引用类型
	//函数后面加（）才是调用

	var callback func(int, int) int
	fmt.Printf("%T ,%#v\n", callback, callback)
	//函数作为一个变量，零值是 nil
	callback = test2
	//将函数赋值给 变量
	rt := callback(1, 4)
	//调用变量--输出值
	fmt.Println(rt)
}
