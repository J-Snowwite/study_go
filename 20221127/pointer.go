package main

import "fmt"

func main() {
	//零值 nil
	var (
		pointerInt *int
		//*int 类型前面加* 就是对应类型的指针
		pointerString *string
	)

	fmt.Printf("%T  %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T  %#v\n", pointerString, pointerString)

	//指针就是特殊的变量，存储一个指向存储空间地址的值
	//赋值
	//使用现有变量，取地址
	age := 32
	fmt.Printf("%T  %#v\n", &age, &age)

	//定义age2
	age2 := age
	fmt.Printf("%T  %#v\n", &age2, &age2)

	//取地址赋值
	pointerInt = &age
	fmt.Println(pointerInt)

	//通过指针获取指针指向的值,通过解引用
	fmt.Println(*pointerInt)

	//通过解引用，对变量进行赋值
	*pointerInt = 33
	fmt.Println(age)

	//age2不变时因为，*pointerInt指向的是age地址存储的值，所以不会影响到age2
	fmt.Println(age2)

	//指针的用途，在一个函数内无法修改函数值的时候，可以通过指针进行修改
	//用指针赋值或者给指针赋值，内存空间不会改变，如果使用变量赋值，会出现不同的地址空间
	age3 := age
	age4 := *pointerInt
	fmt.Println(&age, &age4, &age3, &age2)

	//第二种赋值方式--实际上赋予的是一个空值
	pointerString = new(string)
	fmt.Printf("%#v, %#v \n", pointerString, *pointerString)

	//定义指针的指针
	pp := &pointerString
	fmt.Printf("%T\n", pp)

	//赋值
	**pp = "kk"
	fmt.Println(**pp)
}
