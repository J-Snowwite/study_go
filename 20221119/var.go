package main

import (
	"fmt"
)

//包级别变量
var packageVar string = "package var"

func main() {
	//函数级别变量
	var funcVar string = "func var"

	//函数级别变量--与包级别变量相同
	var packageVar string = "func2"
	{
		//块级别变量
		var blockVar string = "block Var"

		//可以直接打印多个函数
		fmt.Println(packageVar, funcVar, blockVar)

		{ //大括号限制了变量的使用范围，大括号内的只能再内部使用，限制变量的作用域
			//块嵌套级别变量
			var blockVarblock string = "blockVar block"
			fmt.Println(blockVarblock)
		}
	}
	fmt.Println("函数体内", packageVar)
	fmt.Println(funcVar)

	var zeroString string   //定义变量但是不初始化值
	fmt.Println(zeroString) //如果不初始化值，那么就会使用该变量的空值来填充，所以打印出来的是空

	var typeString = "type" //定义变量初始化值，但不指定类型   ，变量会通过值推到变量的类型
	fmt.Println(typeString) //直接打印出值

	shortString := "short"   //使用短声明进行定义，限制：（只能再函数内使用，不能再包级别中使用）
	fmt.Println(shortString) //直接打印值
}
