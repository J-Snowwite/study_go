package main

import "fmt"

func main() {
	name, desc := "AA", "BB"
	func(name string) {
		fmt.Println(name, desc)
		name, desc = "CC", "DD"
		//如果此处师赋值的话，会影响方法外面的变量值
		//如果此处是短声明的话，就不会影响
		fmt.Println(name, desc)
	}("xiaojinteng")
	//函数作用域冲下往上找，找到就停止
	fmt.Println(name, desc)
	//此种方式会导致函数称为一个变量-变量存在回收周期
}
