package main

import (
	"fmt"
)

func test3(flag bool) int {
	if flag {
		return 1
		//函数中一旦碰到return 就不会往下执行逻辑代码了
	}
	fmt.Println("return beforce")
	return 2
}

func main() {
	fmt.Println(test3(true))
	fmt.Println(test3(false))

}
