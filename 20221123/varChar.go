package main

import "fmt"

func main() {
	var msg string = "kk"
	fmt.Printf("%T %s \n", msg, msg)
	//直接打印即可，存在特殊的需要用\ 进行转义
	msg = "我的\n狗"
	fmt.Printf("%T %s", msg, msg)
	//这里就直接转义了，没有换行
	msg = "我的\\n狗"
	fmt.Printf("%T %s\n\n", msg, msg)

	//操作
	var (
		char1 = "abc"
		char2 = "acb"
	)

	//字符连接 +
	fmt.Println(char1 + char2)

	//关系运算 >  <  >=  <= !=  ==
	fmt.Println(char1 > char2)
	fmt.Println(char1 >= char2)

	//赋值运算 = += -=
	char2 += "---kk"
	fmt.Println(char2)

	//索引  切片 ascil
	//如果需要索引的话，一定要保证拿到的一定是ascil可以表示的字符
	fmt.Printf("%T  %#v  %c \n", msg[0], msg[0], msg[0])
	//拿到的第一个字符是uint8 ，中文是无法用uint8 来表示的，所以真正的是拿到了我这个字编码的第一个字节，可以打印出来看一下
	//uint8  0xe6  æ   输出结果中发现并不是  我
	msg = "abcdefgh"
	fmt.Printf("%T  %#v  %c \n", msg[0], msg[0], msg[0])
	//切片的话就是取出其中一定范围的字符
	fmt.Println(msg[1:3])

	// len  计算字节的大小
	fmt.Println(len(msg))

}
