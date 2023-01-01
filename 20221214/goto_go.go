package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println(i)

	goto end
	//会直接跳转到 end 标签处进行执行，可以任意跳，可以往上跳，也可以向下跳
	//i = i + 10
end:
	fmt.Println(i)

	//示例
	index := 0
	total := 0
start:
	index += 1
	total += index
	if index == 100 {
		goto end1
	}
	//goto循环中可以使用goto 跳出循环
	goto start
end1:
	fmt.Println(total)

	//goto 因为违反相对应的逻辑，产生混乱，不建议使用

}
