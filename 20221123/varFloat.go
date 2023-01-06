package main

import "fmt"

func main() {
	var height float32 = 1.68
	var heightType = 1.68
	//默认的浮点数类型是 float64

	fmt.Printf("%T\n%v\n", height, height)
	fmt.Printf("%T\n%v\n", heightType, heightType)

	var k = 1e3
	fmt.Printf("%T\n%v\n", k, k)
	//数字表示可以使用科学计数法

	var (
		f1 = 1.2
		f2 = 2.34
	)
	fmt.Printf("%T  %T  %v  %v\n", f1, f1, f2, f2)
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)
	//四则运算是一样的

	f1++
	f2++
	fmt.Println(f1, f2)
	//自增变量的步长都是1

	//关系运算 > < >= <= == !=
	//如果要判断等于或者不等于，如果差值在一定区间内就是相同
	//浮点数表示并不精切
	fmt.Println(f1 > f2)
	fmt.Println(f1 < f2)
	fmt.Println(f1 >= f2)
	fmt.Println(f1 <= f2)

	//赋值运算 = += -= /= *=
	f1 = 1.1
	f2 = 2.2
	f1 += f2
	//上一行的含义是 f1 = f1 + f2
	fmt.Println(f1)

}
