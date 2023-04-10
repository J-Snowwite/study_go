package main

import "fmt"

var mainVar = "main var"

func mainFunc() {
	fmt.Println("main func")
}

func main() {
	mainFunc()
	fmt.Println(mainVar)

	//1、$env:GO111MODULE="off" 查看 go env |grep MODULE
	//2、编译下运行
	utilsFunc()
	fmt.Println(utilsVar)

	//GOPATH  GOMODULE
	//GO包
	//1、同一个文件夹下所有go文件的包名，必须一致
	//2、关闭GOMODULE
	//GOPATH 在项目目录直接运行 go build 无文件名
	//将当前文件夹下的所有go文件进行编译
	//3、main 包编译为可执行程序
	//3、main 包里面只能有一个main 函数

	//GOPATH--环境变量信息，定义多个目录
	// src:源文件   pkg：编译程序的包文件    bin：编译的可执行程序

}
