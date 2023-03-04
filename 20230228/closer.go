package main

import "fmt"

//返回值是方法，方法和函数是一个东西
//方法 在面向对象里面 就是都叫方法
//方法 在面向过程里面  都叫函数
func addBase(base int) func(int) int {
	//fmt.Println(base)
	//如果只在这个地方使用base，那么输出完成后，base变量地址就会释放
	return func(n int) int {
		return n + base
		//如果在此处使用的话就会导致base不被释放，因为一直在使用
	}
}

//func(n int) int {
//	return n + base
//	//如果在此处使用的话就会导致base不被释放，因为一直在使用
//}
//上述方法如果在函数中被调用，那么可以认为是一个闭包

//变量的生命周期--调用方法时创建，调用完成后销毁
func main() {
	add1 := addBase(1)
	fmt.Println(add1)

	fmt.Println(add1(5))
	//实际返回就是  5 + 1
	//定义一些方法
}
