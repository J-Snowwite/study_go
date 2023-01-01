package main

import "fmt"

func main() {
	multi := [][]string{}
	fmt.Printf("%T,%#v\n\n\n", multi, multi)
	//切片内部可以使用切片或者是数组、字符等等

	multi = append(multi, []string{"A", "B", "C"})
	//对二维切片追加，第一个切片
	multi = append(multi, []string{"D", "E", "f"})
	fmt.Printf("%T ,%#v\n\n\n", multi, multi)

	fmt.Printf("%T,%#v\n\n\n", multi[0], multi[0])
	fmt.Printf("%T,%#v\n\n\n", multi[1], multi[1])
	//打印出两个切片
	fmt.Printf("%T %#v\n\n\n", multi[0][1], multi[0][1])
	fmt.Printf("%T %#v\n\n\n", multi[1][0], multi[1][0])
	//打印出两个切片内容的相对应字符

	multi[0][1] = "xyz"
	multi[1] = append(multi[1], "XXX")
	//在第一个切片内添加一个字符串
	fmt.Println(multi)

}
