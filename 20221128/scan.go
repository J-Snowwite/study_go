package main

import "fmt"

func main() {
	name := ""
	//定义一个空的字符变量
	fmt.Println("请输入你的名字 string类型")
	fmt.Scan(&name)
	//将输入的值，赋值给name这个变量
	//将输入的值，写入到name对应的地址上
	fmt.Printf("你输出的名字是 %#v\n", name)

	age := 0
	//定义一个整形变量
	fmt.Println("请输入你的年龄 int类型")
	fmt.Scan(&age)
	//如果输入值不是int，那么值就会输出0
	//输入值赋值给age
	//输入值写入age 的地址位置
	fmt.Println("你的年龄是", age)

}
