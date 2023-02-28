package main

import "fmt"

func main() {
	name, desc := "AA", "BB"
	func(name string) {
		fmt.Println(name, desc)
	}("xiaojinteng")
	//函数作用域冲下往上找，找到就停止
	fmt.Println(name)
}
