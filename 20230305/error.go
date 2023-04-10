package main

import (
	"errors"
	"fmt"
	"strconv"
)

func div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("除数为0")
	}
	return n1 / n2, nil
}

func main() {
	value, err := strconv.Atoi("XXX")
	//函数返回最后一个参数是错误？
	//go语言的设计哲学
	//希望程序内部如果有错误，通过最后一个返回值显示返回给调用者
	//由调用者决定如何进行处理
	fmt.Printf("%T\n", err)
	fmt.Println(err)
	fmt.Println(value)

	e := fmt.Errorf("自定义错误")
	fmt.Printf("%T %#v\n", e, e)

	e2 := errors.New("自定义错误2")
	fmt.Printf("%T %#v\n", e2, e2)

	if rt, err := div(1, 0); err == nil {
		fmt.Println(rt)
	} else {
		fmt.Println(err)
	}

}
