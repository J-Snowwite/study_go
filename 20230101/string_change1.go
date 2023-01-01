package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	//字节切片  以及  字符串  转换

	ascii := "abc我爱中国"
	fmt.Println([]byte(ascii))
	//转化为字节（字符串）切片
	//	byte  一般用来做二进制处理
	fmt.Println([]rune(ascii))
	//转化为字节     -----      rune 比 byte字节大
	//rune 一般用来做 unicode 处理

	fmt.Println(len(ascii))
	//打印长度----值15 字节长度
	fmt.Println(len([]byte(ascii)))
	//打印长度---值是15--字节长度
	fmt.Println(len([]rune(ascii)))
	//打印长度---值是7---按照unicode字符输出

	fmt.Println(utf8.RuneCountInString(ascii))
	//用utf8 解码形式来统计

	//转换回去
	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))

	//int  float  bool
	fmt.Println(strconv.Itoa('a'))
	//字符转换为  int 转为 ascii
	ch, err := strconv.Atoi("97")
	//ascii 转转为int
	fmt.Println(ch, err)
	fmt.Println(strconv.FormatFloat(3.141592654, 'f', 6, 64))
	//转化为 float 64 位的，也就是小数点后6位
	ch1, err1 := strconv.ParseFloat("3.141592654", 64)
	//不限制位数的64位float
	fmt.Println(ch1, err1)

	fmt.Println(strconv.FormatBool(true))
	//将字符转化为 bool类型的
	fmt.Println(strconv.ParseBool("true"))
	//将字符转化为 bool类型的

	b, err := strconv.ParseInt("6", 10, 8)
	//str 转化为  int  //超过125 就会报错
	fmt.Println(b, err)

	fmt.Println(strconv.FormatInt(15, 2))

}
