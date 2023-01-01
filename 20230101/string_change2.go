package main

import (
	"fmt"
	"strings"
)

func main() {
	//查看官网
	//https://golang.google.cn/pkg/strings/   查看字符串相关
	//https://golang.google.cn/pkg/fmt/   查看fmt模块相关的

	fmt.Println(strings.Compare("a", "b"))
	//比较两个字符  如果相等 输出0 如果小于输出 -1 如果大于输出 1
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("c", "b"))

	//判断第一个参数是否包含第二个参数。是返回true ，不是返回false
	fmt.Println(strings.Contains("我是 AAA", "A"))
	fmt.Println(strings.Contains("我是 AAA", "B"))

}
