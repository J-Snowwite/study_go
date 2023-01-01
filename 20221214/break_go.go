package main

import "fmt"

func main() {

	//break 默认跳出当前for循环，如下所示
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
		for i := 1; i < 3; i++ {
			if i == 2 {
				break
			}
			println("内循环1")
		}
	}

	//break +lables  可以跳出指定循环，2级及以上
OUT:
	for a := 1; a <= 3; a++ {
		fmt.Println(a)
		for b := 1; b < 3; b++ {
			if b == 2 {
				break OUT
			}
			println("内循环2")
		}
	}

	//break 跳出三级循环

shiLi2:
	for c := 1; c <= 3; c++ {
		fmt.Println("循环C")
		for d := 1; d < 3; d++ {
			fmt.Println("循环d")
			for e := 1; e < 3; e++ {
				if e == 2 {
					break shiLi2
				}
				println("循环e")
			}
		}
	}

}
