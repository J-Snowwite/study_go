package main

import "fmt"

//阶乘
// n! =
// n = 0  n! = 1
// n >=1  n! = n * (n-1) * (n-2) *.....* 1
// f(n) = n! = n * f(n - 1)
// 分治 =>  大问题分解为多个相同的小问题 （小问题可以继续拆分，知道某一个可以解决的子问题）
// 递归调用  = >  函数直接或者间接调用自己  （总有一个停止条件）--大部分可以转化为 for 循环

func f(n int64) int64 {
	if n == 0 {
		return 1
	}
	b := 1
	fmt.Println("输出参数", b+1)
	return n * f(n-1)
}

func main() {
	fmt.Println(f(3))

}
