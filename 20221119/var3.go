package main

import "fmt"

func main() {
	var name string = "xjt"
	fmt.Println(name)

	name = "zss" //直接变更值
	fmt.Println(name)
	{ //这里不是定义，是赋值
		name = "AAAAAA"
		fmt.Println(name)
	}
	//因为上述括号中是赋值，所以这里打印的是AAA
	fmt.Println(name)
}
