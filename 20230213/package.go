package main

//导入包  【标准包，自定义包，第三方包】
//包级别的变量  常量  函数
import "fmt"

//定义无惨，无返回值
func sayHello() {
	fmt.Println("hello word")
}

//定义有参，无返回值
func sayHi(name1 string, name2 string) {
	fmt.Println("hi", name1, name2)
}

//有参，有返回值,需要制定返回值类型
func adD(n1 int, n2 int) int {
	return n1 + n2
}

//有参，类型不同，无返回值
func test1(cc int, char1 string) {
	fmt.Println(cc, char1)
}

//有参无返，连续多个相同类型变量
func test2(aa, bb, cc int, dd, ee, ff string) {
	fmt.Println(aa, bb, cc, dd, ee, ff)
}

//可变参数
//1、可变参数，在一个方法中只能有一个
//2、可变参数必须放在函数声明参数列表最后
//func change1(n1 strint,s2 int,c4 ...string),调用前面参数必须写
func test3(args ...string) {
	fmt.Printf("%#v  %T  \n", args, args)

}

func add2(n1, n2 int, arge ...int) int {
	//计算n1 + n2 的值
	total := n1 + n2
	//遍历出arge的所有值求和
	for _, v := range arge {
		total += v
	}
	return total

}

//使用可变函数调用可变函数
func calc(n1, n2 int, args ...int) int {
	//直接调用add 把add 方法结果进行返回
	//args 切片--不能直接调用--必须每一个元素都调用
	//切片可以使用 args...【参数后面加上三个点进行传递】--解操作
	return add2(n1, n2, args...)

}

//多个返回值
func calc2(n1, n2 int) (int, int) {
	a, b := n1, n2
	r1 := add2(a, b)
	r2 := add2(a, b) + 1
	return r1, r2
}

//命名返回值--命名如果使用就全部使用，否则都不用
func calc3(n1, n2 int) (r1, r2 int) {
	//这里无需定义，因为在函数中已经声明了
	r1 = n1 + n2
	r2 = n1 * n2
	return
}

func main() {
	//直接调用sayHello
	//main  调用一个方法 后再调用 sayhello
	sayHello()
	//调用，方法名,传参数数量必须一直
	sayHi("大哥", "二哥")

	//调用，获取返回值,可以不输出
	n1 := adD(4, 5)
	fmt.Println(n1)

	//调用--值类型必须一致
	test1(66, "SIX")

	//相同变量
	test2(11, 22, 33, "a", "b", "c")

	//调用可变参数
	test3()
	test3("AA")
	test3("AA", "BB")

	//调用可变参数的嵌套--切片传递
	n2 := calc(1, 2, 3, 4)
	fmt.Println(n2)

	//直接定义切片调用
	params := []int{4, 5, 6, 7, 7, 8}
	fmt.Println(add2(1, 2, params...))

	//多返回值，调用
	n3, n4 := calc2(3, 6)
	fmt.Println(n3, n4)

	//多返回值调用2
	n5, n6 := calc3(3, 5)
	fmt.Println(n5, n6)
}
