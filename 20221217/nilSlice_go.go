package main

import "fmt"

func main() {
	var nilSlice []int
	var emptySlice []int = []int{} //emptySlice :=[]init{}

	fmt.Printf("%T   %#v\n\n", nilSlice, nilSlice)
	fmt.Printf("%T   %#v\n\n", emptySlice, emptySlice)
	//空切片和nill切片没啥区别，但是这不是一个切片，不相等

	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 1)
	//正常添加元素，输出结果也相同
	fmt.Printf("%#v\n%#v", nilSlice, emptySlice)
	//除了切片类型外，其他的类型此操作会存在差异，例如map，如果是nil类型就无法操作

}
