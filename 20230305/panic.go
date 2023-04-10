package main

import "fmt"

func test() (err error) {
	//recover必须在延迟执行函数内
	//recover捕获异常信息传递给err
	//recover--捕捉第三方库的错误
	//使用recover 后就不会退出程序，而是可以执行下去了
	defer func() {
		fmt.Println("错误信息如下")
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("%s", panicErr)
			//需要将信息转化为err信息
		}
	}()
	fmt.Println("beforce")
	panic("自定义panic")
	//panic显示抛出一个错误
	//panic之后不执行--如果有defer，也会到defer
	fmt.Println("after")
	return
}

func main() {
	fmt.Println("beforce main")
	err := test()
	fmt.Println("after main", err)
}
