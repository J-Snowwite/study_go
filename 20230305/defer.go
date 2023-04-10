package main

import "fmt"

func test() (rt string) {
	//延迟执行，return调用以前，在函数推出之前执行
	defer func() {
		fmt.Println("111")
		rt = "返回错误" //写代码注意，在defer中不要修改返回值
	}()
	defer func() {
		fmt.Println("222")
	}()
	defer func() {
		fmt.Println("333")
	}()
	rt = "test"
	return
}

func test2(n1, n2 int) {
	defer func() {
		//在函数体内，无论是否发生错误，都会执行
		fmt.Println("test2 defer")
	}()
	fmt.Println("before")
	fmt.Println(n1 / n2)
	fmt.Println("after")
}

func main() {
	//defer  函数调用
	//延迟执行，在函数推出之前执行
	//多个defer 会倒序执行--新进后出--堆栈
	defer func() {
		fmt.Println("deferA")
	}()
	defer func() {
		fmt.Println("deferB")
	}()
	defer func() {
		fmt.Println("deferC")
	}()
	fmt.Println("main")

	//调用test
	fmt.Println(test())

	//调用
	test2(3, 0)
}
