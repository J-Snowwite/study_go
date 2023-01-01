package main

import "fmt"

func main() {
	//队列
	//先进先出
	queue := []string{}
	//push 一个消息
	//使用append添加消息
	//append 可以一次性添加多个值，依次添加 append（函数，参数1，参数2......）
	queue = append(queue, "a", "b")
	queue = append(queue, "c")

	//pop
	x := queue[0]
	queue = queue[1:] //切割后的重新赋值
	fmt.Println("1:", x)
	//第一个队列出来
	x = queue[0]
	queue = queue[1:] //切割后的重新赋值
	fmt.Println("2:", x)
	//第二个队列出来
	x = queue[0]
	queue = queue[1:] //切割后的重新赋值
	fmt.Println("3:", x)
	//第三个队列出来

}
