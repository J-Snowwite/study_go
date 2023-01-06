package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		intVal     = 1
		float64Val = 2.5
		stringVal  = "3.3"
	)

	fmt.Println(intVal, float64Val, stringVal)
	//整形转换为浮点型，直接转换即可
	fmt.Printf("%T  %#v\n", float64(intVal), float64(intVal))
	//浮点型转换为整形，直接转换即可,但是会去掉小数点后面的数字
	fmt.Printf("%T  %#v\n", int(float64Val), int(float64Val))

	//整形转化为字符型,直接转换即可--偶尔还是会存在语法报错，不建议使用
	//fmt.Println(string(intVal))

	//字符型不能直接转换为整形,例如下面这句就会报错
	//fmt.Println(int(stringVal))
	//如果需要转换就需要添加一个新的包
	v, err := strconv.Atoi(stringVal)
	fmt.Println(err, v)

	//转化成64为的浮点数
	vv, err := strconv.ParseFloat(stringVal, 64)
	fmt.Println(err, vv)

	//整形转换字符串建议使用strconv来转换
	fmt.Println(strconv.Itoa(32))

	//将float64 转换为float类型，长度是10 ，字节数是64
	fmt.Println(strconv.FormatFloat(float64Val, 'f', 10, 64))
}
