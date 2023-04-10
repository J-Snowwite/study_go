package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("for before ", i)
		//defer,是for 循环结束后执行的
		//打开文件
		//延迟关闭
		//处理（处理出现的错误）
		defer func(i int) {
			fmt.Println("defer", i)
		}(i)
		//循环结束之后才会执行，函数推出时候，i的值为3
		//直接引用传参就可以实现，变量变化
		fmt.Println("for after", i)
	}
}
