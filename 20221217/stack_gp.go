package main

import "fmt"

func main() {
	//堆栈
	//先进后出
	//优先级任务会用到堆栈
	stack := []string{}
	//push
	//append
	stack = append(stack, "a", "b", "c")

	//pop
	//后面的移除,倒数第一个
	x := stack[len(stack)-1:]
	stack = stack[:len(stack)-1]
	fmt.Println("倒数进入消息1", x)
	//后面的移除,倒数第二个
	x = stack[len(stack)-1:]
	stack = stack[:len(stack)-1]
	fmt.Println("倒数进入消息2", x)
	//后面的移除,倒数第三个
	x = stack[len(stack)-1:]
	stack = stack[:len(stack)-1]
	fmt.Println("倒数进入消息3", x)

}
