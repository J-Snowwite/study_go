package main

import "fmt"

func main() {
ConTinue:
	//continue只是跳过，不是跳出，所以基本上不用，跳过后还会继续执行
	for c := 1; c <= 3; c++ {
		fmt.Println("循环C")
		for d := 1; d < 3; d++ {
			fmt.Println("循环d")
			for e := 1; e < 3; e++ {
				if e == 2 {
					continue ConTinue
				}
				println("循环e")
			}
		}
	}
}
