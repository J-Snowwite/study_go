package main

import (
	"fmt"
)

func main() {

	//值类型（存储值）
	//int   float   point     数组   结构体
	//在内存中申请内存新的空间，将age的值拷贝到 tmpAge
	//修改 tmpAge 不影响 age
	//tmpAge  和  age 的在内存空间中存储的是值，所以不影响
	//两个变量的地址是不一样的
	age := 30
	tmpAge := age
	tmpAge = 31
	fmt.Println(age, tmpAge)
	fmt.Printf("%p  %p\n", &age, &tmpAge)

	//引用类型（存储地址）
	//切片  映射   接口
	//申请内存空间--存储地址【也就是指针】
	//在内存中申请新的空间，将 user 的值拷贝到  tmpUser 中
	//修改 tmpUser 中元素的值，是指针指向地址空间的值，但是指针没有变化
	//修改任意指向空间地址的值，访问的目的不变【地址共享】，所以值会变
	//两个变量的地址是一样的
	user := make([]string, 10)
	tmpUser := user
	tmpUser[0] = "SS"
	fmt.Printf("%#v   %#v  \n", user, tmpUser)
	//打印变量的值--%p 打印值
	fmt.Printf("%p   %p\n\n\n\n", user, tmpUser)

	//指针--值类型--指针本身就是值，但是存在一个解引用操作，可以操作其他值
	pointA := &age
	pointB := pointA
	//值操作
	//pointB = &tmpAge
	//fmt.Printf("%#v  %#v  %#v  %#v\n", pointA, pointB, *pointA, *pointB)

	//解引用操作  *  取指针指向地址的值----值相同
	*pointA = 33
	fmt.Printf("%#v  %#v  %#v  %#v\n", pointA, pointB, *pointA, *pointB)

}
