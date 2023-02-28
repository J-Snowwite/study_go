package main

import "fmt"

//值类型，在函数内修改实参的值
//fmt.scan(&age)
//  &取地址 --传参数到变量所在的值中
//  *解引用 ---取指向变量的值
//  go语言中指针不能进行运算，但是可以对指针指向的值进行操作

func change(value int) {
	value += 1
}

func changpointer(pointer *int) {
	*pointer = *pointer + 1
}

func main() {
	value := 1
	change(value)
	fmt.Println(value)
	//输出值任然是 1
	changpointer(&value)
	fmt.Println(value)
	//输出值是2
}
