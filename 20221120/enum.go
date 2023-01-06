package main

import "fmt"

func main() {
	const (
		A = iota // 在常量组内使用 iota初始化值为 0 ，每次调用加1   iota 属于枚举类型
		B
		C
		D = iota //只在第一次调用时候初始化 ，数值还会加一
		E        //打印4
	)
	fmt.Println(A) //打印变量加上换行
	fmt.Println(B)
	fmt.Print(C) //打印变量
	fmt.Print(D)
	fmt.Println("")                   //打印一个空，占位
	fmt.Printf("%T, %v,%#v", E, E, E) //打印出变量的数据类型int,打印出变量E可见的值，按照变量格式打印出，如果字符串就会加上引号
}
