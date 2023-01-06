package main

import "fmt"

func main() {
	var age8 int8 = 31 //字面量，10进制、8进制、16进制
	/*
		8进制：0??  ?<8
		16进制： 0X?? <16  0-9 A-F

		070 => 8进制   56
		078 => 错误输出  应表示 78  10进制
	*/
	var age = 32
	fmt.Printf("%T,%#v,%d\n", age8, age8, age8)
	fmt.Printf("%T.%#v,%d\n", age, age, age)

	//常见操作
	//算数运算  + - * / % ++ --
	a, b := 2, 4
	fmt.Println(a + b) //2
	fmt.Println(a - b) //-2
	fmt.Println(a * b) //8
	fmt.Println(a / b) //0
	//除数不能为 0 ，如果是0 的话编译是不会出现问题，但是运行会出现runtime error
	//编译能通过表示程序的语法没有错误，不代表程序运行是正确的
	fmt.Println(a % b) //2
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	//关系运算
	//关系运算符  >、 <、>=、<=、==、!=
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)

	//位运算
	/*
			7 => 0111
			2 => 0010

		负数：二进制中表示补码，对应整数 +1取反
		2+1=3  -3
		2 ==> 010 取反 ==> 101 ==>3 用负数表示 -3

		按位与 &  ：两个结果都为1， 7&2 = 2  0111与0010，1重叠部分 0010  ==>2
		按位或 |  ：只要一个为1 结果为1  7|2 = 7  0111与0010，1存在部分 0111 ==>2
		取非   ^  ： 二进制 0=>1、 1=>0  010==>101==>-3
		右移位 >> ： 7>>2   111==> 001  =1
		左移位 << :  7<<2   111==> 11100  =28
		and not : &^  取不同位  111 和  010 ，相同为1，不同为0 ==> 101 =5
	*/
	fmt.Println(7 & 2)
	fmt.Println(7 | 2)
	fmt.Println(^2)
	fmt.Println(7 >> 2)
	fmt.Println(7 << 2)
	fmt.Println(7 &^ 2)

	//不同数据类型是无法进行运算的
	var (
		i   int   = 1
		i32 int32 = 1
		i64 int64 = 1
	)
	//需要进行类型转换后才可以进行运算
	//int32、int64、int 只是占用的存储空间是一样大小的，但是不是同一种数据类型
	//type(value)  int32(i)、int(i32)、int64(i)、int(i64)
	fmt.Println(int32(i) + i32)
	fmt.Println(i + int(i64))

	var (
		achar        byte = 'A'
		aint         byte = 64
		unicodePoint rune = '中'
	//字符串存储在内存中主要是01代码，在读取的时候会需要转换
	//转换的时候用到编码（utf8、utf6、gbk）等，每个编码的值对应一个码点，这个码点就是中文或者英文
	)
	fmt.Println(achar, aint, unicodePoint)

	//进制打印---都是int类型，都可以打印出来
	//%d 打印10进制 %b 打印2进制  %o 打印8进制 %x 打印16进制 %U 打印码点 %c 打印字节
	fmt.Printf("%d,%b,%o,%x,%U,%c,%c", achar, 15, 15, 15, unicodePoint, 65, 'A')

}
