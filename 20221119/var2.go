package main

import (
	"fmt"
)

func main() {
	//函数内定义的变量必须使用，否则会报错

	/*
		var name1 string = "one"
		var name2 string = "two"
	*/

	var (
		name1 string = "one"
		name2        = "two"
		desc  string
	)

	fmt.Println(name1, name2, desc)

	/*
		x := "x"
		y := "y"
	*/

	x, y := "x", "y"
	fmt.Println(x, y)

}
