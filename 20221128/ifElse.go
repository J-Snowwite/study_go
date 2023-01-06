package main

import "fmt"

func main() {
	//如果有卖西瓜的，就买一个
	fmt.Println("买10个包子")
	var y string
	fmt.Print("有没有卖西瓜的？\n")
	fmt.Scan(&y)
	fmt.Println("输入的是\n", y)

	if y == "yes" {
		fmt.Println("买一个西瓜")

	} else {
		fmt.Println("只买10个包子")
		if y != "no" {
			fmt.Println("20个包子")
		} else {
			fmt.Print("还是10个包子吧\n\n\n")
		}
	}

	fmt.Println("请输入一个整数")
	var sun1 int = 0
	fmt.Scan(&sun1)
	if sun1 >= 90 {
		fmt.Println("A")
	} else if sun1 >= 70 {
		// else if 可以有 0 到 N 个
		// else  最多只能有一个
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
